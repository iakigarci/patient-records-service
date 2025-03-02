package ports

import (
	"context"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
)

type AuthService interface {
	GenerateToken(user *entities.User) (string, error)
	ValidateToken(tokenString string) (string, error)
}

type UserService interface {
	GetUserByCredentials(ctx context.Context, email, password string) (*entities.User, error)
}

type DiagnosticService interface {
	GetDiagnostics(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error)
}
