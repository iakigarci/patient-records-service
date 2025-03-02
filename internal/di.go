package di

import (
	"github.com/iakigarci/go-ddd-microservice-template/config"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"go.uber.org/zap"
)

// Container holds all the dependencies of the application
type Container struct {
	UserRepository       ports.UserRepository
	DiagnosticRepository ports.DiagnosticRepository
	Config               *config.Config
	Logger               *zap.Logger
}

func NewContainer(cfg *config.Config,
	logger *zap.Logger,
	userRepository ports.UserRepository,
	diagnosticRepository ports.DiagnosticRepository,
) *Container {
	return &Container{
		Config:               cfg,
		Logger:               logger,
		UserRepository:       userRepository,
		DiagnosticRepository: diagnosticRepository,
	}
}

func (c *Container) Shutdown() error {

	return nil
}
