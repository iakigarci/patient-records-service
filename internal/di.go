package di

import (
	"github.com/iakigarci/go-ddd-microservice-template/config"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"go.uber.org/zap"
)

// Container holds all the dependencies of the application
type Container struct {
	UserRepository ports.UserRepository
	Config         *config.Config
	Logger         *zap.Logger
}

func NewContainer(cfg *config.Config,
	logger *zap.Logger,
) *Container {
	return &Container{
		Config: cfg,
		Logger: logger,
	}
}

func (c *Container) Shutdown() error {

	return nil
}
