package handlers

import "github.com/gin-gonic/gin"

type CapsuleHandler struct{}

func NewCapsuleHandler() *CapsuleHandler {
	return &CapsuleHandler{}
}

func (h *CapsuleHandler) Generate(c *gin.Context) {
	c.JSON(201, gin.H{"message": "capsule generation started"})
}

func (h *CapsuleHandler) Get(c *gin.Context) {
	c.JSON(200, gin.H{"message": "capsule details"})
}

func (h *CapsuleHandler) ServeAsset(c *gin.Context) {
	c.JSON(200, gin.H{"message": "asset served"})
}
