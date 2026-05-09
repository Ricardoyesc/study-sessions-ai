package handlers

import "github.com/gin-gonic/gin"

type ColdStartHandler struct{}

func NewColdStartHandler() *ColdStartHandler {
	return &ColdStartHandler{}
}

func (h *ColdStartHandler) Start(c *gin.Context) {
	c.JSON(200, gin.H{"message": "cold start initiated"})
}

func (h *ColdStartHandler) Answer(c *gin.Context) {
	c.JSON(200, gin.H{"message": "answer received"})
}

func (h *ColdStartHandler) Result(c *gin.Context) {
	c.JSON(200, gin.H{"message": "cold start result"})
}
