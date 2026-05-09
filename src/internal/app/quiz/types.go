package quiz

import (
	"sai-server/internal/port"
	a2ui "sai-server/pkg/a2ui"
)

type Engine struct {
	llm     port.LLMClient
	builder *a2ui.Builder
}

type Question struct {
	ID            string
	Topic         string
	Question      string
	Options       []string
	CorrectIndex  int
	Difficulty    float64
}

type Evaluation struct {
	IsCorrect bool
	Feedback  string
	CorrectAnswer string
}

func NewEngine(llm port.LLMClient) *Engine {
	return &Engine{
		llm:     llm,
		builder: a2ui.NewBuilder(),
	}
}
