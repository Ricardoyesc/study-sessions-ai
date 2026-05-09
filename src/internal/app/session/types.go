package session

import (
	"context"
	"math"
	"time"

	"gorm.io/gorm"

	"sai-server/internal/app/dualcode"
	"sai-server/internal/app/quiz"
	"sai-server/internal/app/socratic"
	"sai-server/internal/domain"
)

type State string

const (
	StateColdStart   State = "coldstart"
	StateCapsule     State = "capsule"
	StateQuiz        State = "quiz"
	StateRemediation State = "remediation"
	StateCompleted   State = "completed"
)

type SessionCtx struct {
	Session            *domain.Session
	CurrentState       State
	CurrentQuestion    *quiz.Question
	CurrentRemediation *socratic.Remediation
	CorrectCount       int
	IncorrectCount     int
	TotalItems         int
	CurrentDifficulty  float64
	Progress           float64
}

type Service struct {
	db        *gorm.DB
	dualCode  *dualcode.Orchestrator
	quizEngine *quiz.Engine
	socratic   *socratic.Remediator
}

func NewService(db *gorm.DB, dualCode *dualcode.Orchestrator, quizEngine *quiz.Engine, socratic *socratic.Remediator) *Service {
	return &Service{
		db:         db,
		dualCode:   dualCode,
		quizEngine: quizEngine,
		socratic:   socratic,
	}
}

func (s *Service) CreateSession(ctx context.Context, userID, topic string) (*domain.Session, error) {
	session := domain.Session{
		UserID:            userID,
		State:             string(StateCapsule),
		TargetSuccessRate: 0.85,
		StartedAt:         time.Now(),
	}

	if err := s.db.Create(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *Service) GetSession(ctx context.Context, sessionID string) (*domain.Session, error) {
	session := domain.Session{}
	if err := s.db.Where("id = ?", sessionID).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *Service) LogInteraction(ctx context.Context, sessionID, itype string, payload map[string]interface{}, wasCorrect *bool, responseTimeMs *int) error {
	interaction := domain.Interaction{
		SessionID:      sessionID,
		Type:           itype,
		Payload:        payload,
		WasCorrect:     wasCorrect,
		ResponseTimeMs: responseTimeMs,
	}

	return s.db.Create(&interaction).Error
}

func updateAdaptiveDifficulty(ctx *SessionCtx, wasCorrect bool) {
	ctx.TotalItems++

	if wasCorrect {
		ctx.CorrectCount++
		ctx.CurrentDifficulty = math.Min(ctx.CurrentDifficulty+0.08, 0.95)
	} else {
		ctx.IncorrectCount++
		ctx.CurrentDifficulty = math.Max(ctx.CurrentDifficulty-0.12, 0.10)
	}

	ctx.Progress = float64(ctx.TotalItems) / float64(ctx.TotalItems+3)

	if ctx.TotalItems >= 8 {
		ctx.CurrentState = StateCompleted
	}
}
