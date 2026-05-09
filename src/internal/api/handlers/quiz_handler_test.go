package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"sai-server/internal/api/handlers"
)

func newQuizRouter() *gin.Engine {
	r := gin.New()
	h := handlers.NewQuizHandler()
	r.POST("/api/sessions/:id/quiz/answer", h.Answer)
	return r
}

func TestQuizAnswer_returns200(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/sessions/abc123/quiz/answer", strings.NewReader("{}"))
	req.Header.Set("Content-Type", "application/json")
	newQuizRouter().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status: want 200 got %d — body: %s", w.Code, w.Body.String())
	}
}
