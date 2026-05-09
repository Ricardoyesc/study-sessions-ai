package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"sai-server/internal/api/handlers"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func newUserRouter() *gin.Engine {
	r := gin.New()
	h := handlers.NewUserHandler("test-secret-key")
	r.POST("/api/users/register", h.Register)
	r.POST("/api/users/login", h.Login)
	r.GET("/api/users/me", h.Me)
	return r
}

func TestRegister_validPayload_returns201AndToken(t *testing.T) {
	body, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	newUserRouter().ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("status: want 201 got %d — body: %s", w.Code, w.Body.String())
	}
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if _, ok := resp["token"]; !ok {
		t.Error("response missing 'token' key")
	}
}

func TestRegister_missingEmail_returns400(t *testing.T) {
	body, _ := json.Marshal(map[string]string{"password": "password123"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	newUserRouter().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status: want 400 got %d", w.Code)
	}
}

func TestRegister_invalidEmail_returns400(t *testing.T) {
	body, _ := json.Marshal(map[string]string{"email": "not-an-email", "password": "password123"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	newUserRouter().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status: want 400 got %d", w.Code)
	}
}

func TestRegister_shortPassword_returns400(t *testing.T) {
	body, _ := json.Marshal(map[string]string{"email": "test@example.com", "password": "short"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	newUserRouter().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status: want 400 got %d", w.Code)
	}
}

func TestLogin_validPayload_returns200AndToken(t *testing.T) {
	body, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	newUserRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status: want 200 got %d — body: %s", w.Code, w.Body.String())
	}
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if _, ok := resp["token"]; !ok {
		t.Error("response missing 'token' key")
	}
}

func TestLogin_missingPassword_returns400(t *testing.T) {
	body, _ := json.Marshal(map[string]string{"email": "test@example.com"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	newUserRouter().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status: want 400 got %d", w.Code)
	}
}

func TestMe_returns200WithIdAndEmail(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users/me", nil)
	newUserRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status: want 200 got %d", w.Code)
	}
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if _, ok := resp["id"]; !ok {
		t.Error("response missing 'id' key")
	}
	if _, ok := resp["email"]; !ok {
		t.Error("response missing 'email' key")
	}
}
