package port

import "context"

type LLMClient interface {
	GenerateCompletion(ctx context.Context, prompt string, opts *LLMOptions) (*LLMResponse, error)
	GenerateEmbedding(ctx context.Context, text string) ([]float64, error)
}

type LLMOptions struct {
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
	SystemPrompt string  `json:"system_prompt"`
}

type LLMResponse struct {
	Content      string `json:"content"`
	TokensUsed   int    `json:"tokens_used"`
	Model        string `json:"model"`
	FinishReason string `json:"finish_reason"`
}
