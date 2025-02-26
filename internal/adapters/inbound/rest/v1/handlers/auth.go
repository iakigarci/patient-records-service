package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
)

type AuthHandler struct {
	authService ports.AuthService
}

func New(authService ports.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	token, err := h.authService.Login(c.Request.Context(), email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
