package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"sai-server/internal/api/handlers"
)

func newSocraticRouter() *gin.Engine {
	r := gin.New()
	h := handlers.NewSocraticHandler()
	r.POST("/api/sessions/:id/socratic/response", h.Response)
	return r
}

func TestSocraticResponse_returns200(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/sessions/abc123/socratic/response", strings.NewReader("{}"))
	req.Header.Set("Content-Type", "application/json")
	newSocraticRouter().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: want 200 got %d — body: %s", w.Code, w.Body.String())
	}
}
