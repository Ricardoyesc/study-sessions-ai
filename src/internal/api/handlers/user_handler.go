package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	jwtSecret string
}

func NewUserHandler(jwtSecret string) *UserHandler {
	return &UserHandler{jwtSecret: jwtSecret}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "temp-user-id",
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(h.jwtSecret))

	c.JSON(http.StatusCreated, gin.H{
		"token": tokenStr,
		"user":  gin.H{"email": req.Email},
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "temp-user-id",
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(h.jwtSecret))

	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func (h *UserHandler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":    "temp-user-id",
		"email": "user@example.com",
	})
}
