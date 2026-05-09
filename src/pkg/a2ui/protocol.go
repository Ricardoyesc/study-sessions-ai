package a2ui

import (
	"encoding/json"
	"log/slog"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"sai-server/internal/domain"
)

type Session struct {
	ID        string
	Conn      *websocket.Conn
	Surface   *domain.A2UISurface
	mu        sync.Mutex
	CreatedAt time.Time
}

func (s *Session) SendMessage(msgType MessageType, payload interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	envelope := map[string]interface{}{
		"type":      msgType,
		"payload":   payload,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	return s.Conn.WriteJSON(envelope)
}

func (s *Session) SendError(code, message string) error {
	return s.SendMessage(MsgError, ErrorPayload{Code: code, Message: message})
}

type Engine struct {
	sessions map[string]*Session
	mu       sync.RWMutex
}

func NewEngine() *Engine {
	return &Engine{
		sessions: make(map[string]*Session),
	}
}

func (e *Engine) Register(sessionID string, conn *websocket.Conn) *Session {
	e.mu.Lock()
	defer e.mu.Unlock()

	if existing, ok := e.sessions[sessionID]; ok {
		existing.Conn.Close()
	}

	s := &Session{
		ID:        sessionID,
		Conn:      conn,
		CreatedAt: time.Now(),
	}
	e.sessions[sessionID] = s

	slog.Info("a2ui session registered", "sessionID", sessionID)
	return s
}

func (e *Engine) Unregister(sessionID string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if s, ok := e.sessions[sessionID]; ok {
		s.Conn.Close()
		delete(e.sessions, sessionID)
		slog.Info("a2ui session unregistered", "sessionID", sessionID)
	}
}

func (e *Engine) GetSession(sessionID string) *Session {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.sessions[sessionID]
}

func (e *Engine) SendFull(sessionID string, surface *domain.A2UISurface) error {
	s := e.GetSession(sessionID)
	if s == nil {
		slog.Warn("session not found for send", "sessionID", sessionID)
		return nil
	}

	s.mu.Lock()
	s.Surface = surface
	s.mu.Unlock()

	return s.SendMessage(MsgA2UIFull, surface)
}

func (e *Engine) SendUpdate(sessionID string, updates []DiffUpdate) error {
	s := e.GetSession(sessionID)
	if s == nil {
		return nil
	}

	payload := A2UIUpdatePayload{Updates: updates}

	s.mu.Lock()
	if s.Surface != nil && s.Surface.Components != nil {
		for _, u := range updates {
			if comp, ok := s.Surface.Components[u.ComponentID]; ok {
				for k, v := range u.Props {
					comp.Props[k] = v
				}
				s.Surface.Components[u.ComponentID] = comp
			}
		}
	}
	s.mu.Unlock()

	return s.SendMessage(MsgA2UIUpdate, payload)
}

func (e *Engine) SendDataModelUpdate(sessionID string, update DataModelUpdatePayload) error {
	s := e.GetSession(sessionID)
	if s == nil {
		return nil
	}

	s.mu.Lock()
	if s.Surface != nil {
		dm := &s.Surface.DataModel
		switch update.Path {
		case "theme":
			if v, ok := update.Value.(string); ok {
				dm.Theme = v
			}
		case "fontFamily":
			if v, ok := update.Value.(string); ok {
				dm.FontFamily = v
			}
		case "fontScale":
			if v, ok := update.Value.(float64); ok {
				dm.FontScale = v
			}
		case "colorPalette":
			if v, ok := update.Value.(string); ok {
				dm.ColorPalette = v
			}
		case "highContrast":
			if v, ok := update.Value.(bool); ok {
				dm.HighContrast = v
			}
		case "reducedMotion":
			if v, ok := update.Value.(bool); ok {
				dm.ReducedMotion = v
			}
		case "language":
			if v, ok := update.Value.(string); ok {
				dm.Language = v
			}
		}
	}
	s.mu.Unlock()

	return s.SendMessage(MsgDataModelUpdate, update)
}

func (e *Engine) BroadcastToSession(sessionID string, raw []byte) error {
	s := e.GetSession(sessionID)
	if s == nil {
		return nil
	}

	var msg map[string]interface{}
	if err := json.Unmarshal(raw, &msg); err != nil {
		return err
	}

	slog.Info("a2ui client message", "sessionID", sessionID, "type", msg["type"])
	return nil
}

func (e *Engine) ActiveSessions() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return len(e.sessions)
}

func (e *Engine) CloseAll() {
	e.mu.Lock()
	defer e.mu.Unlock()

	for id, s := range e.sessions {
		s.Conn.Close()
		delete(e.sessions, id)
	}
}
