package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"sai-server/internal/app/session"
)

type QuizHandler struct {
	pipeline *session.Pipeline
}

func NewQuizHandler(pipeline *session.Pipeline) *QuizHandler {
	return &QuizHandler{pipeline: pipeline}
}

func (h *QuizHandler) Answer(c *gin.Context) {
	sessionID := c.Param("id")

	var req struct {
		SelectedIndex int `json:"selected_index" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.pipeline.EvaluateAnswer(sessionID, req.SelectedIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_ = strconv.Itoa(req.SelectedIndex) // can be used for logging

	c.JSON(http.StatusOK, gin.H{
		"state":          result.State,
		"is_correct":     result.IsCorrect,
		"correct_answer": result.CorrectAnswer,
		"feedback":       result.Feedback,
		"message":        result.Message,
	})
}
