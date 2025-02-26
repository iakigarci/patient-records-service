package http_gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iakigarci/go-ddd-microservice-template/config"
	di "github.com/iakigarci/go-ddd-microservice-template/internal"
	"github.com/iakigarci/go-ddd-microservice-template/internal/adapters/inbound/rest/v1/handlers"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/services/auth"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/services/user"
)

type Router struct {
	Router    *gin.Engine
	container *di.Container
}

func New(config *config.Config, container *di.Container) *Router {
	r := gin.Default()
	router := &Router{
		Router:    r,
		container: container,
	}

	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	v1 := r.Group("/v1")
	{
		router.indexRoutes(v1)
		router.authRoutes(v1)
	}

	r.Run(fmt.Sprintf(":%d", config.HTTP.Port))

	return router
}

func (r *Router) indexRoutes(rg *gin.RouterGroup) {
	indexRoutes := rg.Group("/")
	{
		indexRoutes.GET("/health", handlers.HealthCheck)
	}
}

func (r *Router) authRoutes(rg *gin.RouterGroup) {
	userService := user.New(
		user.WithUserRepository(r.container.UserRepository),
		user.WithLogger(r.container.Logger),
	)
	authService := auth.New(
		userService,
		[]byte("secret"),
		r.container.Logger,
	)

	authHandler := handlers.NewAuthHandler(authService, userService)

	authRoutes := rg.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
	}
}
