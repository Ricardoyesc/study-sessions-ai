package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"sai-server/internal/api/handlers"
)

func newSessionRouter() *gin.Engine {
	r := gin.New()
	h := handlers.NewSessionHandler()
	r.POST("/api/sessions", h.Create)
	r.GET("/api/sessions/:id", h.Get)
	r.GET("/api/sessions/:id/next", h.Next)
	r.POST("/api/sessions/:id/a11y", h.UpdateAccessibility)
	return r
}

func TestSessionCreate_returns201(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/sessions", strings.NewReader("{}"))
	req.Header.Set("Content-Type", "application/json")
	newSessionRouter().ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("status: want 201 got %d — body: %s", w.Code, w.Body.String())
	}
}

func TestSessionGet_returns200(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/sessions/abc123", nil)
	newSessionRouter().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: want 200 got %d", w.Code)
	}
}

func TestSessionNext_returns200(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/sessions/abc123/next", nil)
	newSessionRouter().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: want 200 got %d", w.Code)
	}
}

func TestSessionUpdateAccessibility_returns200(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/sessions/abc123/a11y", strings.NewReader("{}"))
	req.Header.Set("Content-Type", "application/json")
	newSessionRouter().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: want 200 got %d", w.Code)
	}
}
