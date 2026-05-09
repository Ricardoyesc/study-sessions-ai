package llm

import (
	"context"

	"github.com/sashabaranov/go-openai"

	"sai-server/internal/port"
)

type OpenAI struct {
	client *openai.Client
	model  string
}

func NewOpenAI(apiKey, model string) *OpenAI {
	return &OpenAI{
		client: openai.NewClient(apiKey),
		model:  model,
	}
}

func (o *OpenAI) GenerateCompletion(ctx context.Context, prompt string, opts *port.LLMOptions) (*port.LLMResponse, error) {
	systemPrompt := ""
	maxTokens := 1024
	temperature := 0.7

	if opts != nil {
		systemPrompt = opts.SystemPrompt
		if opts.MaxTokens > 0 {
			maxTokens = opts.MaxTokens
		}
		if opts.Temperature > 0 {
			temperature = opts.Temperature
		}
	}

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: prompt},
	}
	if systemPrompt != "" {
		messages = append([]openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
		}, messages...)
	}

	resp, err := o.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       o.model,
		Messages:    messages,
		MaxTokens:   maxTokens,
		Temperature: float32(temperature),
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return &port.LLMResponse{Content: "No response generated"}, nil
	}

	return &port.LLMResponse{
		Content:      resp.Choices[0].Message.Content,
		TokensUsed:   resp.Usage.TotalTokens,
		Model:        resp.Model,
		FinishReason: string(resp.Choices[0].FinishReason),
	}, nil
}

func (o *OpenAI) GenerateEmbedding(ctx context.Context, text string) ([]float64, error) {
	resp, err := o.client.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Model: openai.AdaEmbeddingV2,
		Input: text,
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, nil
	}

	embedding := make([]float64, len(resp.Data[0].Embedding))
	for i, v := range resp.Data[0].Embedding {
		embedding[i] = float64(v)
	}
	return embedding, nil
}
