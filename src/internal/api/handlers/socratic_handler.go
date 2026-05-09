package handlers

import "github.com/gin-gonic/gin"

type SocraticHandler struct{}

func NewSocraticHandler() *SocraticHandler {
	return &SocraticHandler{}
}

func (h *SocraticHandler) Response(c *gin.Context) {
	c.JSON(200, gin.H{"message": "socratic response received"})
}
