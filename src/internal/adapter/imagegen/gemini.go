package imagegen

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"sai-server/internal/port"
)

type Gemini struct {
	apiKey string
	model  string
	client *http.Client
}

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
	GenerationConfig *geminiGenerationConfig `json:"generationConfig,omitempty"`
}

type geminiContent struct {
	Parts []geminiPart `json:"parts"`
}

type geminiPart struct {
	Text string `json:"text,omitempty"`
}

type geminiGenerationConfig struct {
	ResponseModalities []string `json:"responseModalities"`
}

type geminiResponse struct {
	Candidates []geminiCandidate `json:"candidates"`
	Error      *geminiError      `json:"error,omitempty"`
}

type geminiCandidate struct {
	Content *geminiResponseContent `json:"content,omitempty"`
}

type geminiResponseContent struct {
	Parts []geminiResponsePart `json:"parts"`
}

type geminiResponsePart struct {
	Text       string             `json:"text,omitempty"`
	InlineData *geminiInlineData  `json:"inlineData,omitempty"`
}

type geminiInlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
}

type geminiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewGemini(apiKey, model string) *Gemini {
	return &Gemini{
		apiKey: apiKey,
		model:  model,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

func (g *Gemini) GenerateImage(ctx context.Context, prompt string, opts *port.ImageGenOptions) (*port.ImageGenResponse, error) {
	imagePrompt := "Generate an image: " + prompt
	if opts != nil && opts.Style != "" {
		imagePrompt = fmt.Sprintf("Generate an image in %s style: %s", opts.Style, prompt)
	}

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		g.model, g.apiKey,
	)

	reqBody := geminiRequest{
		Contents: []geminiContent{
			{
				Parts: []geminiPart{
					{Text: imagePrompt},
				},
			},
		},
		GenerationConfig: &geminiGenerationConfig{
			ResponseModalities: []string{"IMAGE", "TEXT"},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal gemini request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create gemini request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	slog.Debug("gemini image gen request", "model", g.model, "prompt_len", len(prompt))

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("gemini request failed: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read gemini response: %w", err)
	}

	var geminiResp geminiResponse
	if err := json.Unmarshal(respBytes, &geminiResp); err != nil {
		return nil, fmt.Errorf("unmarshal gemini response: %w", err)
	}

	if geminiResp.Error != nil {
		return nil, fmt.Errorf("gemini api error [%d]: %s", geminiResp.Error.Code, geminiResp.Error.Message)
	}

	if len(geminiResp.Candidates) == 0 || geminiResp.Candidates[0].Content == nil {
		return nil, fmt.Errorf("gemini returned no image candidates")
	}

	var urls []string
	var revisedPrompt string

	for _, part := range geminiResp.Candidates[0].Content.Parts {
		if part.InlineData != nil && part.InlineData.Data != "" {
			dataURL := fmt.Sprintf("data:%s;base64,%s", part.InlineData.MimeType, part.InlineData.Data)
			urls = append(urls, dataURL)

			imageBytes, decodeErr := base64.StdEncoding.DecodeString(part.InlineData.Data)
			if decodeErr == nil {
				slog.Debug("gemini image generated", "size_bytes", len(imageBytes), "mime", part.InlineData.MimeType)
			}
		}
		if part.Text != "" && revisedPrompt == "" {
			revisedPrompt = part.Text
		}
	}

	if len(urls) == 0 {
		return nil, fmt.Errorf("gemini response contained no image data")
	}

	return &port.ImageGenResponse{
		ImageURLs:     urls,
		RevisedPrompt: revisedPrompt,
	}, nil
}
