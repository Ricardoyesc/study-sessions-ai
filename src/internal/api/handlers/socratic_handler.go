package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sai-server/internal/app/session"
)

type SocraticHandler struct {
	pipeline *session.Pipeline
}

func NewSocraticHandler(pipeline *session.Pipeline) *SocraticHandler {
	return &SocraticHandler{pipeline: pipeline}
}

func (h *SocraticHandler) Response(c *gin.Context) {
	sessionID := c.Param("id")

	var req struct {
		StudentResponse string `json:"student_response" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.pipeline.ProcessRemediationResponse(sessionID, req.StudentResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"state":   result.State,
		"message": result.Message,
	})
}
