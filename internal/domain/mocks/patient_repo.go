package mocks

import (
	"context"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
)

type MockPatientRepository struct {
	GetPatientByIDFn func(ctx context.Context, id string) (*entities.Patient, error)
}

func (m *MockPatientRepository) GetPatientByID(ctx context.Context, id string) (*entities.Patient, error) {
	return m.GetPatientByIDFn(ctx, id)
}
