package diagnostic

import (
	"context"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"go.uber.org/zap"
)

type DiagnosticService struct {
	diagnosticRepository ports.DiagnosticRepository
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
