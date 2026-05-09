package dualcode

import (
	"context"
	"fmt"
	"log/slog"

	"sai-server/internal/domain"
	"sai-server/internal/port"
	a2ui "sai-server/pkg/a2ui"
)

func (o *Orchestrator) GenerateCapsule(ctx context.Context, topic string) (*Capsule, error) {
	slog.Info("generating capsule", "topic", topic)

	textResp, err := o.llm.GenerateCompletion(ctx, buildTextPrompt(topic), &port.LLMOptions{
		MaxTokens:   800,
		Temperature: 0.7,
		SystemPrompt: "Eres un tutor experto. Genera contenido educativo claro y conciso en español. Usa Markdown para el formato.",
	})
	if err != nil {
		slog.Warn("llm text generation failed, using fallback", "error", err)
		textResp = fallbackText(topic)
	}

	imagePrompt := fmt.Sprintf(
		"Diagrama educativo sobre %s, estilo clean textbook illustration, labeled, professional, no text in image",
		topic,
	)
	imgResp, imgErr := o.imageGen.GenerateImage(ctx, imagePrompt, &port.ImageGenOptions{
		Width:  1024,
		Height: 1024,
	})
	if imgErr != nil {
		slog.Warn("image generation failed", "error", imgErr)
	}

	var imageURL, imageAlt string
	if imgResp != nil && len(imgResp.ImageURLs) > 0 {
		imageURL = imgResp.ImageURLs[0]
		imageAlt = fmt.Sprintf("Diagrama sobre %s", topic)
	} else {
		imageURL = ""
		imageAlt = ""
	}

	components := map[string]domain.A2UIComponent{
		"root":            a2ui.NewColumn("root", []string{"header", "body"}, nil),
		"header":          a2ui.NewRow("header", []string{"title"}, map[string]interface{}{"alignment": "center"}),
		"title":           a2ui.NewText("title", topic, "h2"),
		"body":            a2ui.NewCard("body", []string{"text-content", "image-diagram"}, nil),
		"text-content":    a2ui.NewRichText("text-content", textResp.Content),
	}

	if imageURL != "" {
		components["image-diagram"] = a2ui.NewImage("image-diagram", imageURL, imageAlt)
	} else {
		components["image-diagram"] = a2ui.NewText("image-diagram", "[Imagen no disponible — configura GEMINI_API_KEY]", "caption")
	}

	surface := o.builder.BuildCapsuleSurface(fmt.Sprintf("capsule-%s", topic), topic, components, a2ui.DefaultDataModel())

	return &Capsule{
		Topic:       topic,
		Text:        textResp.Content,
		ImageURL:    imageURL,
		ImageAlt:    imageAlt,
		A2UISurface: surface,
	}, nil
}

func buildTextPrompt(topic string) string {
	return fmt.Sprintf(`Genera una lección educativa sobre "%s" en español. Incluye:
1. Una introducción breve (2-3 frases)
2. Conceptos clave con viñetas (3-5 puntos)
3. Un ejemplo práctico o aplicación
4. Un dato interesante o curiosidad

Usa formato Markdown (## para títulos, ** para negrita, - para viñetas).`, topic)
}

func fallbackText(topic string) *port.LLMResponse {
	return &port.LLMResponse{
		Content: fmt.Sprintf(`## %s

### Introducción
Este es contenido educativo sobre **%s**. El sistema está funcionando en modo fallback —
configura `+"`OPENAI_API_KEY`"+` para generar contenido real con IA.

### Conceptos Clave
- Concepto principal de %s
- Aplicación práctica en el mundo real
- Relación con otros temas de estudio

### Ejemplo
Un ejemplo práctico de %s se puede observar en situaciones cotidianas.

### Dato curioso
¿Sabías que %s tiene aplicaciones fascinantes en múltiples disciplinas?

---
*Modo demostración — contenido pre-generado*`, topic, topic, topic, topic, topic),
	}
}
