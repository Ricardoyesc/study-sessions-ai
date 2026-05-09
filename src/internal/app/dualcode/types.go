package dualcode

import (
	"sai-server/internal/domain"
	"sai-server/internal/port"
	a2ui "sai-server/pkg/a2ui"
)

type Orchestrator struct {
	llm      port.LLMClient
	imageGen port.ImageGenClient
	tts      port.TTSClient
	builder  *a2ui.Builder
}

type Capsule struct {
	ID         string
	Topic      string
	Text       string
	ImageURL   string
	ImageAlt   string
	AudioURL   string
	A2UISurface *domain.A2UISurface
}

func NewOrchestrator(llm port.LLMClient, imageGen port.ImageGenClient, tts port.TTSClient) *Orchestrator {
	return &Orchestrator{
		llm:      llm,
		imageGen: imageGen,
		tts:      tts,
		builder:  a2ui.NewBuilder(),
	}
}
