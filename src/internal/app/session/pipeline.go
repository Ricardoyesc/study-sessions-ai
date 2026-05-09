package session

import (
	"context"
	"fmt"
	"log/slog"

	"sai-server/internal/app/quiz"
	"sai-server/internal/domain"
)

type Pipeline struct {
	service  *Service
	sessions map[string]*SessionCtx
}

func NewPipeline(service *Service) *Pipeline {
	return &Pipeline{
		service:  service,
		sessions: make(map[string]*SessionCtx),
	}
}

func (p *Pipeline) GetOrCreateCtx(sessionID string) *SessionCtx {
	if ctx, ok := p.sessions[sessionID]; ok {
		return ctx
	}

	session, err := p.service.GetSession(context.Background(), sessionID)
	if err != nil {
		slog.Error("failed to get session", "sessionID", sessionID, "error", err)
		return nil
	}

	sctx := &SessionCtx{
		Session:           session,
		CurrentState:      State(session.State),
		CurrentDifficulty: 0.50,
	}
	p.sessions[sessionID] = sctx
	return sctx
}

func (p *Pipeline) NextItem(sessionID, topic string) (*PipelineResult, error) {
	ctx := p.GetOrCreateCtx(sessionID)
	if ctx == nil {
		return nil, fmt.Errorf("session not found: %s", sessionID)
	}

	slog.Info("pipeline next item", "sessionID", sessionID, "state", ctx.CurrentState, "difficulty", ctx.CurrentDifficulty, "progress", ctx.Progress)

	switch ctx.CurrentState {
	case StateCapsule, StateColdStart:
		return p.generateCapsule(ctx, topic)
	case StateQuiz:
		return p.generateQuiz(ctx, topic)
	case StateRemediation:
		return p.generateRemediation(ctx, topic)
	case StateCompleted:
		return &PipelineResult{State: StateCompleted, Message: "Sesión completada. ¡Buen trabajo!"}, nil
	default:
		ctx.CurrentState = StateCapsule
		return p.generateCapsule(ctx, topic)
	}
}

func (p *Pipeline) EvaluateAnswer(sessionID string, selectedIndex int) (*PipelineResult, error) {
	ctx := p.GetOrCreateCtx(sessionID)
	if ctx == nil || ctx.CurrentQuestion == nil {
		return nil, fmt.Errorf("no active question for session %s", sessionID)
	}

	eval := p.service.quizEngine.EvaluateAnswer(ctx.CurrentQuestion, selectedIndex)

	wasCorrect := eval.IsCorrect
	updateAdaptiveDifficulty(ctx, wasCorrect)

	p.service.LogInteraction(context.Background(), sessionID, "quiz_answer", map[string]interface{}{
		"question_id":      ctx.CurrentQuestion.ID,
		"selected_index":   selectedIndex,
		"correct_index":    ctx.CurrentQuestion.CorrectIndex,
		"topic":            ctx.CurrentQuestion.Topic,
		"difficulty":       ctx.CurrentDifficulty,
		"correct_count":    ctx.CorrectCount,
		"incorrect_count":  ctx.IncorrectCount,
		"total_items":      ctx.TotalItems,
	}, &wasCorrect, nil)

	if eval.IsCorrect {
		ctx.CurrentState = StateCapsule
		ctx.CurrentQuestion = nil
		return &PipelineResult{
			State:         StateCapsule,
			Message:       fmt.Sprintf("¡Correcto! (%d/%d aciertos)", ctx.CorrectCount, ctx.TotalItems),
			IsCorrect:     true,
			CorrectAnswer: eval.CorrectAnswer,
			Feedback:      eval.Feedback,
		}, nil
	}

	ctx.CurrentState = StateRemediation
	return &PipelineResult{
		State:         StateRemediation,
		Message:       fmt.Sprintf("Incorrecto. Revisemos este concepto. (%d/%d aciertos)", ctx.CorrectCount, ctx.TotalItems),
		IsCorrect:     false,
		CorrectAnswer: eval.CorrectAnswer,
		Feedback:      eval.Feedback,
	}, nil
}

func (p *Pipeline) ProcessRemediationResponse(sessionID, studentResponse string) (*PipelineResult, error) {
	ctx := p.GetOrCreateCtx(sessionID)
	if ctx == nil {
		return nil, fmt.Errorf("session not found: %s", sessionID)
	}

	p.service.LogInteraction(context.Background(), sessionID, "socratic_response", map[string]interface{}{
		"response": studentResponse,
	}, nil, nil)

	ctx.CurrentState = StateCapsule
	ctx.CurrentRemediation = nil

	return &PipelineResult{
		State:   StateCapsule,
		Message: "Reflexión recibida. Continuemos con el siguiente tema.",
	}, nil
}

func (p *Pipeline) generateCapsule(ctx *SessionCtx, topic string) (*PipelineResult, error) {
	capsule, err := p.service.dualCode.GenerateCapsule(context.Background(), topic)
	if err != nil {
		return nil, err
	}

	ctx.CurrentState = StateQuiz
	ctx.Session.State = string(StateQuiz)
	p.service.db.Save(ctx.Session)

	p.service.LogInteraction(context.Background(), ctx.Session.ID, "capsule_generated", map[string]interface{}{
		"topic": topic,
	}, nil, nil)

	return &PipelineResult{
		State:       StateCapsule,
		Topic:       capsule.Topic,
		A2UISurface: capsule.A2UISurface,
		Message:     fmt.Sprintf("Cápsula generada: %s (dificultad: %.0f%%)", topic, ctx.CurrentDifficulty*100),
	}, nil
}

func (p *Pipeline) generateQuiz(ctx *SessionCtx, topic string) (*PipelineResult, error) {
	question, err := p.service.quizEngine.GenerateQuestion(context.Background(), topic, ctx.CurrentDifficulty)
	if err != nil {
		return nil, err
	}

	ctx.CurrentQuestion = question

	components := map[string]domain.A2UIComponent{
		"quiz-root": {
			ID:       "quiz-root",
			Type:     "Column",
			Children: []string{"quiz-progress", "quiz-card"},
			Props:    map[string]interface{}{"gap": 16, "padding": 24},
		},
		"quiz-progress": {
			ID:   "quiz-progress",
			Type: "ProgressBar",
			Props: map[string]interface{}{
				"value": ctx.Progress,
				"max":   1.0,
			},
		},
		"quiz-card": {
			ID:   "quiz-card",
			Type: "QuizCard",
			Props: map[string]interface{}{
				"question": question.Question,
				"options":  question.Options,
				"mode":     "single_choice",
			},
			Events: map[string]string{
				"onSubmit": fmt.Sprintf("/api/sessions/%s/quiz/answer", ctx.Session.ID),
			},
		},
	}

	surface := &domain.A2UISurface{
		SurfaceID:     ctx.Session.ID,
		RootComponent: "quiz-root",
		Components:    components,
		DataModel:     domain.A2UIDataModel{},
	}

	return &PipelineResult{
		State:       StateQuiz,
		Question:    question,
		A2UISurface: surface,
		Message:     fmt.Sprintf("Pregunta sobre: %s (dificultad: %.0f%%)", topic, ctx.CurrentDifficulty*100),
	}, nil
}

func (p *Pipeline) generateRemediation(ctx *SessionCtx, topic string) (*PipelineResult, error) {
	wrongAnswer := ""
	correctAnswer := ""
	if ctx.CurrentQuestion != nil {
		wrongAnswer = ctx.CurrentQuestion.Options[ctx.CurrentQuestion.CorrectIndex]
		correctAnswer = ctx.CurrentQuestion.Options[ctx.CurrentQuestion.CorrectIndex]
	}

	remediation, err := p.service.socratic.GenerateRemediation(context.Background(), topic, wrongAnswer, correctAnswer)
	if err != nil {
		return nil, err
	}

	ctx.CurrentRemediation = remediation

	return &PipelineResult{
		State:       StateRemediation,
		A2UISurface: remediation.A2UISurface,
		Message:     "Revisemos este concepto con más detalle.",
	}, nil
}

type PipelineResult struct {
	State         State
	Topic         string
	Question      *quiz.Question
	A2UISurface   *domain.A2UISurface
	Message       string
	IsCorrect     bool
	CorrectAnswer string
	Feedback      string
}
