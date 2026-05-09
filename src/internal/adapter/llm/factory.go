package llm

import (
	"context"

	"sai-server/internal/port"
)

func NewOpenAIClient(apiKey, model string) port.LLMClient {
	if apiKey == "" {
		return &noopLLM{model: model}
	}
	return NewOpenAI(apiKey, model)
}

type noopLLM struct {
	model string
}

func (n *noopLLM) GenerateCompletion(_ context.Context, _ string, _ *port.LLMOptions) (*port.LLMResponse, error) {
	return &port.LLMResponse{
		Content:    "This is a fallback response. Configure OPENAI_API_KEY for real completions.",
		TokensUsed: 0,
		Model:      n.model,
	}, nil
}

func (n *noopLLM) GenerateEmbedding(_ context.Context, _ string) ([]float64, error) {
	return make([]float64, 1536), nil
}
