package tts

import (
	"context"

	"sai-server/internal/port"
)

func NewTTSClient(apiKey string, provider string) port.TTSClient {
	if apiKey == "" {
		return &noopTTS{}
	}
	switch provider {
	case "openai":
		return NewOpenAI(apiKey, "tts-1")
	default:
		return NewOpenAI(apiKey, "tts-1")
	}
}

type noopTTS struct{}

func (n *noopTTS) SynthesizeSpeech(_ context.Context, _ string, _ *port.TTSOptions) (*port.TTSResponse, error) {
	return &port.TTSResponse{
		AudioData:  []byte{},
		Format:     "mp3",
		DurationMs: 0,
	}, nil
}
