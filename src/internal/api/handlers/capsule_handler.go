package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sai-server/internal/app/dualcode"
)

type CapsuleHandler struct {
	orch *dualcode.Orchestrator
}

func NewCapsuleHandler(orch *dualcode.Orchestrator) *CapsuleHandler {
	return &CapsuleHandler{orch: orch}
}

func (h *CapsuleHandler) Generate(c *gin.Context) {
	var req struct {
		Topic string `json:"topic" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	capsule, err := h.orch.GenerateCapsule(c.Request.Context(), req.Topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"topic":   capsule.Topic,
		"text":    capsule.Text,
		"image_url": capsule.ImageURL,
		"surface": capsule.A2UISurface,
	})
}

func (h *CapsuleHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "capsule details"})
}

func (h *CapsuleHandler) ServeAsset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "asset served"})
}
