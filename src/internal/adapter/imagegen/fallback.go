package imagegen

import (
	"context"
	"log/slog"

	"sai-server/internal/port"
)

type geminiWithFallback struct {
	primary   *Gemini
	fallback  *DALLE
}

func NewGeminiWithFallback(gemini *Gemini, dalle *DALLE) port.ImageGenClient {
	return &geminiWithFallback{
		primary:  gemini,
		fallback: dalle,
	}
}

func (g *geminiWithFallback) GenerateImage(ctx context.Context, prompt string, opts *port.ImageGenOptions) (*port.ImageGenResponse, error) {
	resp, err := g.primary.GenerateImage(ctx, prompt, opts)
	if err == nil {
		slog.Info("image generated via gemini", "urls", len(resp.ImageURLs))
		return resp, nil
	}

	slog.Warn("gemini image generation failed, falling back to DALL·E", "error", err)

	return g.fallback.GenerateImage(ctx, prompt, opts)
}
