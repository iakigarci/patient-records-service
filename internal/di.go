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
	PatientRepository    ports.PatientRepository
	Config               *config.Config
	Logger               *zap.Logger
}

func NewContainer(cfg *config.Config,
	logger *zap.Logger,
	userRepository ports.UserRepository,
	diagnosticRepository ports.DiagnosticRepository,
	patientRepository ports.PatientRepository,
) *Container {
	return &Container{
		Config:               cfg,
		Logger:               logger,
		UserRepository:       userRepository,
		DiagnosticRepository: diagnosticRepository,
		PatientRepository:    patientRepository,
	}
}

func (c *Container) Shutdown() error {
	return nil
}
