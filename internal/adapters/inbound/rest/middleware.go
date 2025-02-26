package http_gin

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iakigarci/go-ddd-microservice-template/internal/adapters/inbound/rest/v1/handlers"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func AuthMiddleware(authService ports.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, handlers.ErrorResponse{Error: "Authorization header required"})
			c.Abort()
			return
		}

		email, err := authService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, handlers.ErrorResponse{Error: "Invalid token"})
			c.Abort()
			return
		}

		c.Set("email", email)
		c.Next()
	}
}
