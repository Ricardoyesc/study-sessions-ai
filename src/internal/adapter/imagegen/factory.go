package imagegen

import (
	"context"

	"sai-server/internal/port"
)

func NewImageGenClient(geminiKey, openAIKey, geminiModel string) port.ImageGenClient {
	hasGemini := geminiKey != ""
	hasOpenAI := openAIKey != ""

	if hasGemini && hasOpenAI {
		gemini := NewGemini(geminiKey, geminiModel)
		dalle := NewDALLE(openAIKey)
		return NewGeminiWithFallback(gemini, dalle)
	}

	if hasGemini {
		return NewGemini(geminiKey, geminiModel)
	}

	if hasOpenAI {
		return NewDALLE(openAIKey)
	}

	return &noopImageGen{}
}

type noopImageGen struct{}

func (n *noopImageGen) GenerateImage(_ context.Context, _ string, _ *port.ImageGenOptions) (*port.ImageGenResponse, error) {
	return &port.ImageGenResponse{
		ImageURLs: []string{""},
	}, nil
}
