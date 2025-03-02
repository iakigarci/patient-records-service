package mocks

import (
	"context"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
)

type MockDiagnosticRepository struct {
	CreateDiagnosticFn func(ctx context.Context, diagnostic *entities.Diagnostic) error
	GetDiagnosticsFn   func(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error)
}

func (m *MockDiagnosticRepository) GetDiagnostics(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error) {
	return m.GetDiagnosticsFn(ctx, filter)
}

func (m *MockDiagnosticRepository) CreateDiagnostic(ctx context.Context, diagnostic *entities.Diagnostic) error {
	return m.CreateDiagnosticFn(ctx, diagnostic)
}
