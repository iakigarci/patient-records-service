package user

import (
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"go.uber.org/zap"
)

type UserBuilderOption func(*UserSvc)

func New(opts ...UserBuilderOption) ports.UserService {
	options := &UserSvc{}
	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithUserRepository(userRepository ports.UserRepository) UserBuilderOption {
	return func(opts *UserSvc) {
		opts.userRepository = userRepository
	}
}

func WithLogger(logger *zap.Logger) UserBuilderOption {
	return func(opts *UserSvc) {
		opts.logger = logger
	}
}
