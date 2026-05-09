package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"sai-server/internal/port"
)

type llmQuizResponse struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Correct  int      `json:"correct"`
}

func (e *Engine) GenerateQuestion(ctx context.Context, topic string, difficulty float64) (*Question, error) {
	slog.Info("generating quiz question", "topic", topic, "difficulty", difficulty)

	resp, err := e.llm.GenerateCompletion(ctx, buildQuizPrompt(topic, difficulty), &port.LLMOptions{
		MaxTokens:   300,
		Temperature: 0.8,
		SystemPrompt: "Eres un generador de preguntas educativas. Responde SOLO en el formato JSON especificado, sin markdown ni texto adicional.",
	})
	if err != nil {
		slog.Warn("quiz generation failed, using fallback", "error", err)
		return fallbackQuestion(topic), nil
	}

	q := parseQuestion(resp.Content, topic)
	return q, nil
}

func (e *Engine) EvaluateAnswer(question *Question, selectedIndex int) *Evaluation {
	correct := selectedIndex == question.CorrectIndex
	eval := &Evaluation{
		IsCorrect:     correct,
		CorrectAnswer: question.Options[question.CorrectIndex],
	}

	if correct {
		eval.Feedback = "¡Correcto! Buen trabajo."
	} else {
		eval.Feedback = fmt.Sprintf("Incorrecto. La respuesta correcta era: %s", question.Options[question.CorrectIndex])
	}

	return eval
}

func buildQuizPrompt(topic string, difficulty float64) string {
	diffLabel := "básica"
	if difficulty > 0.5 {
		diffLabel = "intermedia"
	}
	if difficulty > 0.8 {
		diffLabel = "avanzada"
	}

	return fmt.Sprintf(`Genera una pregunta de opción múltiple de dificultad %s sobre "%s".

Responde EXACTAMENTE en este formato JSON (sin markdown, sin comillas adicionales):
{
  "question": "texto de la pregunta",
  "options": ["A) op1", "B) op2", "C) op3", "D) op4"],
  "correct": 0
}

Donde "correct" es el índice (0-3) de la respuesta correcta.`, diffLabel, topic)
}

func parseQuestion(raw, topic string) *Question {
	cleaned := strings.TrimSpace(raw)
	cleaned = strings.TrimPrefix(cleaned, "```json")
	cleaned = strings.TrimPrefix(cleaned, "```")
	cleaned = strings.TrimSuffix(cleaned, "```")
	cleaned = strings.TrimSpace(cleaned)

	var parsed llmQuizResponse
	if err := json.Unmarshal([]byte(cleaned), &parsed); err != nil {
		slog.Warn("failed to parse quiz JSON, using fallback", "raw", raw[:min(100, len(raw))], "error", err)
		return fallbackQuestion(topic)
	}

	if parsed.Question == "" || len(parsed.Options) < 2 {
		slog.Warn("quiz JSON incomplete, using fallback", "question", parsed.Question, "options", len(parsed.Options))
		return fallbackQuestion(topic)
	}

	if parsed.Correct < 0 || parsed.Correct >= len(parsed.Options) {
		parsed.Correct = 0
	}

	return &Question{
		ID:           fmt.Sprintf("q-%s-%d", strings.ReplaceAll(topic, " ", "-"), len(topic)),
		Topic:        topic,
		Question:     parsed.Question,
		Options:      parsed.Options,
		CorrectIndex: parsed.Correct,
		Difficulty:   0.5,
	}
}

func fallbackQuestion(topic string) *Question {
	return &Question{
		ID:           fmt.Sprintf("q-fallback-%s", topic),
		Topic:        topic,
		Question:     fmt.Sprintf("¿Cuál es el concepto principal de %s?", topic),
		Options:      []string{"A) La definición básica", "B) Su aplicación práctica", "C) Su historia y origen", "D) Su relación con otros conceptos"},
		CorrectIndex: 0,
		Difficulty:   0.5,
	}
}
