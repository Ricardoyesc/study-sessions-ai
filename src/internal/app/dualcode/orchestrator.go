package dualcode

import (
	"context"
	"encoding/base64"
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
	}

	var audioURL string
	audioText := fmt.Sprintf("Lección sobre %s. %s", topic, extractPlainText(textResp.Content))
	if len(audioText) > 1000 {
		audioText = audioText[:1000]
	}
	audioResp, audioErr := o.tts.SynthesizeSpeech(ctx, audioText, &port.TTSOptions{
		Voice: "nova",
		Speed: 1.0,
	})
	if audioErr != nil {
		slog.Warn("audio generation failed", "error", audioErr)
	}
	if audioResp != nil && len(audioResp.AudioData) > 0 {
		audioURL = "data:audio/mp3;base64," + base64.StdEncoding.EncodeToString(audioResp.AudioData)
	}

	children := []string{"text-content", "image-diagram"}
	if audioURL != "" {
		children = append(children, "audio-player")
	}

	components := map[string]domain.A2UIComponent{
		"root":         a2ui.NewColumn("root", []string{"header", "progress", "body"}, map[string]interface{}{"gap": 16, "padding": 24}),
		"header":       a2ui.NewRow("header", []string{"title"}, map[string]interface{}{"alignment": "center"}),
		"title":        a2ui.NewText("title", topic, "h2"),
		"progress":     a2ui.NewProgressBar("progress", 0.25, 1.0),
		"body":         a2ui.NewCard("body", children, map[string]interface{}{"elevation": 2}),
		"text-content": a2ui.NewRichText("text-content", textResp.Content),
	}

	if imageURL != "" {
		components["image-diagram"] = a2ui.NewImage("image-diagram", imageURL, imageAlt)
	} else {
		components["image-diagram"] = a2ui.NewText("image-diagram", "[Imagen no disponible — configura GEMINI_API_KEY y OPENAI_API_KEY]", "caption")
	}

	if audioURL != "" {
		components["audio-player"] = a2ui.NewAudioPlayer("audio-player", audioURL)
	}

	dm := a2ui.DefaultDataModel()
	dm.Language = "es"

	surface := o.builder.BuildCapsuleSurface(fmt.Sprintf("capsule-%s", topic), topic, components, dm)

	return &Capsule{
		Topic:       topic,
		Text:        textResp.Content,
		ImageURL:    imageURL,
		ImageAlt:    imageAlt,
		AudioURL:    audioURL,
		A2UISurface: surface,
	}, nil
}

func extractPlainText(md string) string {
	result := ""
	for _, c := range md {
		if c == '#' || c == '*' || c == '-' || c == '_' || c == '`' {
			continue
		}
		result += string(c)
	}
	if len(result) > 500 {
		result = result[:500]
	}
	return result
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
