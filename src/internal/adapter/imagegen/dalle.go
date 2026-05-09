package imagegen

import (
	"context"

	"github.com/sashabaranov/go-openai"

	"sai-server/internal/port"
)

type DALLE struct {
	client *openai.Client
}

func NewDALLE(apiKey string) *DALLE {
	return &DALLE{
		client: openai.NewClient(apiKey),
	}
}

func (d *DALLE) GenerateImage(ctx context.Context, prompt string, opts *port.ImageGenOptions) (*port.ImageGenResponse, error) {
	size := "1024x1024"
	n := 1

	if opts != nil {
		if opts.NumImages > 0 {
			n = opts.NumImages
		}
		switch {
		case opts.Width == 512 && opts.Height == 512:
			size = "512x512"
		case opts.Width == 256 && opts.Height == 256:
			size = "256x256"
		case opts.Width == 1792 && opts.Height == 1024:
			size = "1792x1024"
		case opts.Width == 1024 && opts.Height == 1792:
			size = "1024x1792"
		}
	}

	resp, err := d.client.CreateImage(ctx, openai.ImageRequest{
		Model:   openai.CreateImageModelDallE3,
		Prompt:  prompt,
		N:       n,
		Size:    size,
		Quality: openai.CreateImageQualityStandard,
	})
	if err != nil {
		return nil, err
	}

	urls := make([]string, len(resp.Data))
	for i, img := range resp.Data {
		urls[i] = img.URL
	}

	return &port.ImageGenResponse{
		ImageURLs:     urls,
		RevisedPrompt: resp.Data[0].RevisedPrompt,
	}, nil
}
