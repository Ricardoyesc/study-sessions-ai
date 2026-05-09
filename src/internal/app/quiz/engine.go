package quiz

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"sai-server/internal/port"
)

func (e *Engine) GenerateQuestion(ctx context.Context, topic string, difficulty float64) (*Question, error) {
	slog.Info("generating quiz question", "topic", topic, "difficulty", difficulty)

	resp, err := e.llm.GenerateCompletion(ctx, buildQuizPrompt(topic, difficulty), &port.LLMOptions{
		MaxTokens:   300,
		Temperature: 0.8,
		SystemPrompt: "Eres un generador de preguntas educativas. Responde SOLO en el formato JSON especificado, sin texto adicional.",
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
	return &Question{
		ID:           fmt.Sprintf("q-%s-%d", strings.ReplaceAll(topic, " ", "-"), len(topic)),
		Topic:        topic,
		Question:     "¿Cuál es el concepto principal de " + topic + "?",
		Options:      []string{"A) Opción 1", "B) Opción 2", "C) Opción 3", "D) Opción 4"},
		CorrectIndex: 0,
		Difficulty:   0.5,
	}
}

func fallbackQuestion(topic string) *Question {
	return &Question{
		ID:           fmt.Sprintf("q-fallback-%s", topic),
		Topic:        topic,
		Question:     fmt.Sprintf("¿Cuál es el concepto principal de %s?", topic),
		Options:      []string{"A) Primera opción", "B) Segunda opción", "C) Tercera opción", "D) Cuarta opción"},
		CorrectIndex: 0,
		Difficulty:   0.5,
	}
}
