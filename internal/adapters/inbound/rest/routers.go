package http_gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iakigarci/go-ddd-microservice-template/config"
	di "github.com/iakigarci/go-ddd-microservice-template/internal"
	"github.com/iakigarci/go-ddd-microservice-template/internal/adapters/inbound/rest/v1/handlers"
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

	routeV1(r)

	r.Run(fmt.Sprintf(":%d", config.HTTP.Port))

	return router
}

func routeV1(r *gin.Engine) {
	indexRoutes := r.Group("/")
	{
		indexRoutes.GET("/health", handlers.HealthCheck)
	}

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", handlers.Login)
	}

	diagnosticRoutes := r.Group("/diagnostic")
	{
		diagnosticRoutes.GET("/", handlers.GetDiagnostics)
		diagnosticRoutes.POST("/", handlers.CreateDiagnostic)
	}
}
