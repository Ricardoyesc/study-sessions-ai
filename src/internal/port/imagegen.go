package port

import "context"

type ImageGenClient interface {
	GenerateImage(ctx context.Context, prompt string, opts *ImageGenOptions) (*ImageGenResponse, error)
}

type ImageGenOptions struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Style     string `json:"style"`
	NumImages int    `json:"num_images"`
}

type ImageGenResponse struct {
	ImageURLs []string `json:"image_urls"`
	RevisedPrompt string `json:"revised_prompt,omitempty"`
}
