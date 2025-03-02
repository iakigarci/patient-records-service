package diagnostic

import (
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"go.uber.org/zap"
)

type ServiceOption func(*DiagnosticService)

func New(opts ...ServiceOption) *DiagnosticService {
	options := &DiagnosticService{}
	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithRepository(repo ports.DiagnosticRepository) ServiceOption {
	return func(s *DiagnosticService) {
		s.diagnosticRepository = repo
	}
}

func WithLogger(logger *zap.Logger) ServiceOption {
	return func(s *DiagnosticService) {
		s.logger = logger
	}
}
