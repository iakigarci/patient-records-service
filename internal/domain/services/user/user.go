package user

import (
	"context"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserSvc struct {
	userRepository ports.UserRepository
	logger         *zap.Logger
}

func (svc *UserSvc) GetUserByCredentials(ctx context.Context, email, password string) (*entities.User, error) {
	user, err := svc.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		svc.logger.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}

	svc.logger.Info("user", zap.Any("user", user))

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		svc.logger.Error("failed to compare password hash", zap.Error(err))
		return nil, err
	}

	return user, nil
}
