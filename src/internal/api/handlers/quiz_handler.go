package handlers

import "github.com/gin-gonic/gin"

type QuizHandler struct{}

func NewQuizHandler() *QuizHandler {
	return &QuizHandler{}
}

func (h *QuizHandler) Answer(c *gin.Context) {
	c.JSON(200, gin.H{"message": "quiz answer received"})
}
