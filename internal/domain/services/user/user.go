package user

import (
	"context"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"go.uber.org/zap"
)

type UserSvc struct {
	userRepository ports.UserRepository
	logger         *zap.Logger
}

func (svc *UserSvc) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	return svc.userRepository.GetUserByEmail(ctx, email)
}
