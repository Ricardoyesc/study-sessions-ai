package handlers

import "github.com/gin-gonic/gin"

type SessionHandler struct{}

func NewSessionHandler() *SessionHandler {
	return &SessionHandler{}
}

func (h *SessionHandler) Create(c *gin.Context) {
	c.JSON(201, gin.H{"message": "session created"})
}

func (h *SessionHandler) Get(c *gin.Context) {
	c.JSON(200, gin.H{"message": "session details"})
}

func (h *SessionHandler) Next(c *gin.Context) {
	c.JSON(200, gin.H{"message": "next item"})
}

func (h *SessionHandler) UpdateAccessibility(c *gin.Context) {
	c.JSON(200, gin.H{"message": "accessibility updated"})
}
