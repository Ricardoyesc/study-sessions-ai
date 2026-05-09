package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sai-server/internal/app/session"
)

type SessionHandler struct {
	svc      *session.Service
	pipeline *session.Pipeline
}

func NewSessionHandler(svc *session.Service, pipeline *session.Pipeline) *SessionHandler {
	return &SessionHandler{svc: svc, pipeline: pipeline}
}

func (h *SessionHandler) Create(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}

	var req struct {
		Topic string `json:"topic" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := h.svc.CreateSession(c.Request.Context(), userIDVal.(string), req.Topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":     s.ID,
		"state":  s.State,
		"topic":  req.Topic,
		"ws_url": "ws://127.0.0.1:8080/ws/session/" + s.ID,
	})
}

func (h *SessionHandler) Get(c *gin.Context) {
	sessionID := c.Param("id")

	s, err := h.svc.GetSession(c.Request.Context(), sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":                  s.ID,
		"state":               s.State,
		"target_success_rate": s.TargetSuccessRate,
		"started_at":          s.StartedAt,
	})
}

func (h *SessionHandler) Next(c *gin.Context) {
	sessionID := c.Param("id")
	topic := c.Query("topic")

	if topic == "" {
		topic = "Conceptos generales"
	}

	result, err := h.pipeline.NextItem(sessionID, topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"state":   result.State,
		"topic":   result.Topic,
		"message": result.Message,
		"surface": result.A2UISurface,
	})
}

func (h *SessionHandler) UpdateAccessibility(c *gin.Context) {
	sessionID := c.Param("id")

	var req struct {
		FontFamily    string  `json:"fontFamily"`
		FontScale     float64 `json:"fontScale"`
		HighContrast  bool    `json:"highContrast"`
		ReducedMotion bool    `json:"reducedMotion"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": sessionID,
		"message":    "accessibility preferences saved",
	})
}
