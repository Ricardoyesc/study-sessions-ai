package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"sai-server/internal/api/handlers"
)

func newCapsuleRouter() *gin.Engine {
	r := gin.New()
	h := handlers.NewCapsuleHandler()
	r.POST("/api/capsules/generate", h.Generate)
	r.GET("/api/capsules/:id", h.Get)
	r.GET("/api/assets/:type/:filename", h.ServeAsset)
	return r
}

func TestCapsuleGenerate_returns201(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/capsules/generate", strings.NewReader("{}"))
	req.Header.Set("Content-Type", "application/json")
	newCapsuleRouter().ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("status: want 201 got %d — body: %s", w.Code, w.Body.String())
	}
}

func TestCapsuleGet_returns200(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/capsules/cap1", nil)
	newCapsuleRouter().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: want 200 got %d", w.Code)
	}
}

func TestCapsuleServeAsset_returns200(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/assets/audio/file.mp3", nil)
	newCapsuleRouter().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: want 200 got %d", w.Code)
	}
}
