package main

import (
	"log"

	"github.com/iakigarci/go-ddd-microservice-template/config"
	di "github.com/iakigarci/go-ddd-microservice-template/internal"
	http_gin "github.com/iakigarci/go-ddd-microservice-template/internal/adapters/inbound/rest"
	"github.com/iakigarci/go-ddd-microservice-template/internal/adapters/outbound/postgres"
	httpserver "github.com/iakigarci/go-ddd-microservice-template/pkg/http"
	"github.com/iakigarci/go-ddd-microservice-template/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig[config.Config]()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := logger.New(cfg)

	container := getDIContainer(cfg, logger)
	httpServer := startServers(cfg, container)

	if err := <-httpServer.Notify(); err != nil {
		logger.Error("Failed to start server", zap.Error(err))
	}

	shutdown(httpServer, container, logger)
}

func getDIContainer(cfg *config.Config, logger *zap.Logger) *di.Container {
	dbClient, err := postgres.NewClient(cfg, logger)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepository := postgres.NewUserRepository(dbClient.DB)
	diagnosticRepository := postgres.NewDiagnosticRepository(dbClient.DB)
	patientRepository := postgres.NewPatientRepository(dbClient.DB)

	return di.NewContainer(cfg,
		logger,
		userRepository,
		diagnosticRepository,
		patientRepository,
	)
}

func startServers(cfg *config.Config, container *di.Container) *httpserver.Server {
	httpServer := http_gin.New(cfg, container)

	server := httpserver.New(cfg, httpServer.Router)
	return server
}

func shutdown(server *httpserver.Server, container *di.Container, log *zap.Logger) {
	if shutdownErr := server.Shutdown(); shutdownErr != nil {
		log.Error("httpServer.Shutdown", zap.Error(shutdownErr))
	}

	if err := container.Shutdown(); err != nil {
		log.Error("container.Shutdown", zap.Error(err))
	}
}
