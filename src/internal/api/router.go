package api

import (
	"github.com/gin-gonic/gin"

	"sai-server/internal/api/handlers"
)

func RegisterRoutes(r *gin.Engine, h *handlers.Handlers, jwtSecret string) {
	api := r.Group("/api")
	{
		api.POST("/users/register", h.User.Register)
		api.POST("/users/login", h.User.Login)

		auth := api.Group("")
		auth.Use(func(c *gin.Context) { c.Next() })
		{
			auth.GET("/users/me", h.User.Me)

			auth.POST("/coldstart/start", h.ColdStart.Start)
			auth.POST("/coldstart/answer", h.ColdStart.Answer)
			auth.GET("/coldstart/:id/result", h.ColdStart.Result)

			auth.POST("/sessions", h.Session.Create)
			auth.GET("/sessions/:id", h.Session.Get)
			auth.GET("/sessions/:id/next", h.Session.Next)
			auth.POST("/sessions/:id/quiz/answer", h.Quiz.Answer)
			auth.POST("/sessions/:id/socratic/response", h.Socratic.Response)
			auth.POST("/sessions/:id/a11y", h.Session.UpdateAccessibility)

			auth.POST("/capsules/generate", h.Capsule.Generate)
			auth.GET("/capsules/:id", h.Capsule.Get)

			auth.GET("/assets/:type/:filename", h.Capsule.ServeAsset)
		}
	}

	ws := r.Group("/ws")
	{
		ws.GET("/session/:sessionId", h.A2UI.HandleWebSocket)
	}

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/readyz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ready"})
	})
}
