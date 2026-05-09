package tts

import (
	"context"
	"io"

	"github.com/sashabaranov/go-openai"

	"sai-server/internal/port"
)

type OpenAITTS struct {
	client *openai.Client
	model  string
}

func NewOpenAI(apiKey, model string) *OpenAITTS {
	return &OpenAITTS{
		client: openai.NewClient(apiKey),
		model:  model,
	}
}

func (t *OpenAITTS) SynthesizeSpeech(ctx context.Context, text string, opts *port.TTSOptions) (*port.TTSResponse, error) {
	voice := openai.VoiceAlloy
	speed := 1.0

	if opts != nil {
		if opts.Speed > 0 {
			speed = opts.Speed
		}
		switch opts.Voice {
		case "nova":
			voice = openai.VoiceNova
		case "echo":
			voice = openai.VoiceEcho
		case "fable":
			voice = openai.VoiceFable
		case "onyx":
			voice = openai.VoiceOnyx
		case "shimmer":
			voice = openai.VoiceShimmer
		}
	}

	resp, err := t.client.CreateSpeech(ctx, openai.CreateSpeechRequest{
		Model: openai.SpeechModel(t.model),
		Input: text,
		Voice: voice,
		Speed: speed,
	})
	if err != nil {
		return nil, err
	}

	defer resp.Close()
	audioData, err := io.ReadAll(resp)
	if err != nil {
		return nil, err
	}

	return &port.TTSResponse{
		AudioData:  audioData,
		Format:     "mp3",
		DurationMs: len(audioData) / 16,
	}, nil
}
