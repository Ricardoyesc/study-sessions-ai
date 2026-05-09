package socratic

import (
	"context"
	"fmt"
	"log/slog"

	"sai-server/internal/domain"
	"sai-server/internal/port"
	a2ui "sai-server/pkg/a2ui"
)

type Remediator struct {
	llm     port.LLMClient
	builder *a2ui.Builder
}

type Remediation struct {
	FeynmanExplanation string
	Analogy            string
	ElaborativeQuestion string
	Topic              string
	A2UISurface        *domain.A2UISurface
}

func NewRemediator(llm port.LLMClient) *Remediator {
	return &Remediator{
		llm:     llm,
		builder: a2ui.NewBuilder(),
	}
}

func (r *Remediator) GenerateRemediation(ctx context.Context, topic, wrongAnswer, correctAnswer string) (*Remediation, error) {
	slog.Info("generating socratic remediation", "topic", topic)

	resp, err := r.llm.GenerateCompletion(ctx, buildRemediationPrompt(topic, wrongAnswer, correctAnswer), &port.LLMOptions{
		MaxTokens:   600,
		Temperature: 0.7,
		SystemPrompt: "Eres un tutor socrático experto. Explica conceptos de forma simple usando la Técnica Feynman y analogías cotidianas. Responde en español.",
	})
	if err != nil {
		slog.Warn("remediation generation failed, using fallback", "error", err)
		resp = fallbackRemediation(topic, correctAnswer)
	}

	components := map[string]domain.A2UIComponent{
		"root":              a2ui.NewColumn("root", []string{"explanation", "analogy", "socratic-prompt"}, nil),
		"explanation":       a2ui.NewCard("explanation", []string{"exp-title", "exp-text"}, map[string]interface{}{"elevation": 1}),
		"exp-title":         a2ui.NewText("exp-title", "Explicación (Técnica Feynman)", "h3"),
		"exp-text":          a2ui.NewRichText("exp-text", resp.Content),
		"analogy":           a2ui.NewCard("analogy", []string{"analogy-title", "analogy-text"}, map[string]interface{}{"elevation": 1}),
		"analogy-title":    a2ui.NewText("analogy-title", "Analogía", "h3"),
		"analogy-text":     a2ui.NewRichText("analogy-text", fmt.Sprintf("Piensa en %s como si fuera...", topic)),
		"socratic-prompt":   a2ui.NewSocraticDialog(
			"socratic-prompt",
			fmt.Sprintf("¿Por qué crees que tu respuesta '%s' no era correcta? Explica con tus propias palabras qué entendiste de la explicación anterior.", wrongAnswer),
			topic,
			fmt.Sprintf("/api/sessions/current/socratic/response"),
		),
	}

	surface := r.builder.BuildSocraticSurface(fmt.Sprintf("remediation-%s", topic), components, a2ui.DefaultDataModel())

	return &Remediation{
		FeynmanExplanation: resp.Content,
		Topic:              topic,
		A2UISurface:        surface,
	}, nil
}

func buildRemediationPrompt(topic, wrongAnswer, correctAnswer string) string {
	return fmt.Sprintf(`El estudiante respondió incorrectamente a una pregunta sobre "%s".
Respuesta del estudiante: "%s"
Respuesta correcta: "%s"

Aplica la Técnica Feynman:
1. Explica el concepto correcto en los términos más simples posibles
2. Usa una analogía de la vida cotidiana
3. Identifica por qué la respuesta del estudiante era incorrecta
4. Termina con una pregunta para que el estudiante reflexione (Interrogación Elaborativa)

Responde en español, en formato Markdown.`, topic, wrongAnswer, correctAnswer)
}

func fallbackRemediation(topic, correctAnswer string) *port.LLMResponse {
	return &port.LLMResponse{
		Content: fmt.Sprintf(`## Explicación: %s

### En términos simples
%s es un concepto que se puede entender descomponiéndolo en partes más pequeñas.

### ¿Por qué tu respuesta era incorrecta?
La respuesta correcta era **%s**. Es importante revisar este concepto porque es fundamental para entender temas más avanzados.

### Pregunta de reflexión
¿Puedes explicar con tus propias palabras por qué la respuesta correcta tiene sentido?

---
*Modo demostración — configura `+"`OPENAI_API_KEY`"+` para generar explicaciones personalizadas*`, topic, topic, correctAnswer),
	}
}
