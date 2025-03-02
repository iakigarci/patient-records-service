package diagnostic

import (
	"context"
	"fmt"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"go.uber.org/zap"
)

type DiagnosticService struct {
	diagnosticRepository ports.DiagnosticRepository
	patientRepository    ports.PatientRepository
	logger               *zap.Logger
}

func (s *DiagnosticService) GetDiagnostics(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error) {
	diagnostics, err := s.diagnosticRepository.GetDiagnostics(ctx, filter)
	if err != nil {
		s.logger.Error("failed to get diagnostics", zap.Error(err))
		return nil, err
	}

	return diagnostics, nil
}

func (s *DiagnosticService) CreateDiagnostic(ctx context.Context, diagnostic *entities.Diagnostic) error {
	patient, err := s.patientRepository.GetPatientByID(ctx, diagnostic.PatientID)
	if err != nil {
		s.logger.Error("failed to get patient", zap.Error(err))
		return fmt.Errorf("patient not found: %v", err)
	}

	if patient == nil {
		err = fmt.Errorf("patient not found with ID: %s", diagnostic.PatientID)
		s.logger.Error("failed to get patient", zap.Error(err))
		return err
	}

	if err := s.diagnosticRepository.CreateDiagnostic(ctx, diagnostic); err != nil {
		s.logger.Error("failed to create diagnostic", zap.Error(err))
		return err
	}
	return nil
}
