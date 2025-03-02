package ports

import (
	"context"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
}

type DiagnosticRepository interface {
	GetDiagnostics(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error)
}
