package port

import "context"

type TTSClient interface {
	SynthesizeSpeech(ctx context.Context, text string, opts *TTSOptions) (*TTSResponse, error)
}

type TTSOptions struct {
	Language string  `json:"language"`
	Voice    string  `json:"voice"`
	Speed    float64 `json:"speed"`
}

type TTSResponse struct {
	AudioURL    string `json:"audio_url"`
	AudioData   []byte `json:"audio_data,omitempty"`
	DurationMs  int    `json:"duration_ms"`
	Format      string `json:"format"`
}
