package handlers

import "github.com/gin-gonic/gin"

type ColdStartHandler struct{}

func NewColdStartHandler() *ColdStartHandler {
	return &ColdStartHandler{}
}

func (h *ColdStartHandler) Start(c *gin.Context) {
	items := []gin.H{
		{"id": "cs-1", "question": "¿Cuál es tu nivel de familiaridad con el método científico?", "difficulty": 0.3},
		{"id": "cs-2", "question": "¿Puedes explicar qué es una variable independiente?", "difficulty": 0.5},
		{"id": "cs-3", "question": "¿Qué es un grupo de control en un experimento?", "difficulty": 0.7},
	}
	c.JSON(200, gin.H{
		"diagnostic_id": "diag-001",
		"items":         items,
		"total_items":   len(items),
	})
}

func (h *ColdStartHandler) Answer(c *gin.Context) {
	var req struct {
		DiagnosticID string `json:"diagnostic_id" binding:"required"`
		ItemID       string `json:"item_id" binding:"required"`
		Answer       string `json:"answer" binding:"required"`
		ResponseTimeMs int  `json:"response_time_ms"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	theta := 0.0
	uncertainty := 1.0

	switch req.ItemID {
	case "cs-1":
		theta = 0.2
	case "cs-2":
		theta = 0.5
	case "cs-3":
		theta = 0.8
	}

	c.JSON(200, gin.H{
		"estimated_theta": theta,
		"uncertainty":     uncertainty,
		"message":        "Answer recorded",
	})
}

func (h *ColdStartHandler) Result(c *gin.Context) {
	cluster := "intermediate"

	c.JSON(200, gin.H{
		"diagnostic_id":    c.Param("id"),
		"estimated_theta":  0.55,
		"theta_uncertainty": 0.25,
		"cluster":          cluster,
		"confidence":       0.85,
	})
}
