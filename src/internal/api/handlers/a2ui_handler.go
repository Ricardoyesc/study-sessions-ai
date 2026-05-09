package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	a2ui "sai-server/pkg/a2ui"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 4096
)

type A2UIHandler struct {
	upgrader websocket.Upgrader
	engine   *a2ui.Engine
}

func NewA2UIHandler(engine *a2ui.Engine) *A2UIHandler {
	return &A2UIHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		engine: engine,
	}
}

func (h *A2UIHandler) HandleWebSocket(c *gin.Context) {
	sessionID := c.Param("sessionId")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sessionId required"})
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		slog.Error("websocket upgrade failed", "error", err)
		return
	}

	session := h.engine.Register(sessionID, conn)
	defer h.engine.Unregister(sessionID)

	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	slog.Info("websocket connected", "sessionID", sessionID)

	go h.writePump(session, conn)

	for {
		_, raw, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				slog.Warn("websocket read error", "sessionID", sessionID, "error", err)
			}
			break
		}

		if err := h.engine.BroadcastToSession(sessionID, raw); err != nil {
			slog.Warn("failed to process message", "sessionID", sessionID, "error", err)
		}
	}
}

func (h *A2UIHandler) writePump(session *a2ui.Session, conn *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for range ticker.C {
		if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			slog.Warn("ping failed", "sessionID", session.ID, "error", err)
			return
		}
	}
}

func (h *A2UIHandler) Engine() *a2ui.Engine {
	return h.engine
}
