package diagnostic

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/mocks"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestDiagnosticService_CreateDiagnostic(t *testing.T) {
	logger := zap.NewNop()

	tests := []struct {
		name          string
		diagnostic    *entities.Diagnostic
		setupMocks    func(dr *mocks.MockDiagnosticRepository, pr *mocks.MockPatientRepository)
		expectedError string
		shouldSucceed bool
	}{
		{
			name: "successful creation",
			diagnostic: &entities.Diagnostic{
				PatientID: "123",
				Diagnosis: "Common cold",
				Date:      time.Now(),
			},
			setupMocks: func(dr *mocks.MockDiagnosticRepository, pr *mocks.MockPatientRepository) {
				pr.GetPatientByIDFn = func(ctx context.Context, id string) (*entities.Patient, error) {
					return &entities.Patient{ID: "123", Name: "John Doe"}, nil
				}
				dr.CreateDiagnosticFn = func(ctx context.Context, diagnostic *entities.Diagnostic) error {
					return nil
				}
			},
			shouldSucceed: true,
		},
		{
			name: "patient not found",
			diagnostic: &entities.Diagnostic{
				PatientID: "0f6f6fce-127e-4e83-b865-e78bcbab881d",
				Diagnosis: "Fever",
				Date:      time.Now(),
			},
			setupMocks: func(dr *mocks.MockDiagnosticRepository, pr *mocks.MockPatientRepository) {
				pr.GetPatientByIDFn = func(ctx context.Context, id string) (*entities.Patient, error) {
					return nil, nil
				}
			},
			expectedError: "patient not found with ID: 0f6f6fce-127e-4e83-b865-e78bcbab881d",
			shouldSucceed: false,
		},
		{
			name: "error getting patient",
			diagnostic: &entities.Diagnostic{
				PatientID: "0f6f6fce-127e-4e83-b865-e78bcbab881d",
				Diagnosis: "Headache",
				Date:      time.Now(),
			},
			setupMocks: func(dr *mocks.MockDiagnosticRepository, pr *mocks.MockPatientRepository) {
				pr.GetPatientByIDFn = func(ctx context.Context, id string) (*entities.Patient, error) {
					return nil, fmt.Errorf("database error")
				}
			},
			expectedError: "patient not found: database error",
			shouldSucceed: false,
		},
		{
			name: "error creating diagnostic",
			diagnostic: &entities.Diagnostic{
				PatientID: "0f6f6fce-127e-4e83-b865-e78bcbab881d",
				Diagnosis: "Flu",
				Date:      time.Now(),
			},
			setupMocks: func(dr *mocks.MockDiagnosticRepository, pr *mocks.MockPatientRepository) {
				pr.GetPatientByIDFn = func(ctx context.Context, id string) (*entities.Patient, error) {
					return &entities.Patient{ID: "0f6f6fce-127e-4e83-b865-e78bcbab881d", Name: "John Doe"}, nil
				}
				dr.CreateDiagnosticFn = func(ctx context.Context, diagnostic *entities.Diagnostic) error {
					return fmt.Errorf("failed to create diagnostic")
				}
			},
			expectedError: "failed to create diagnostic",
			shouldSucceed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDiagRepo := &mocks.MockDiagnosticRepository{}
			mockPatientRepo := &mocks.MockPatientRepository{}
			tt.setupMocks(mockDiagRepo, mockPatientRepo)

			service := &DiagnosticService{
				diagnosticRepository: mockDiagRepo,
				patientRepository:    mockPatientRepo,
				logger:               logger,
			}

			err := service.CreateDiagnostic(context.Background(), tt.diagnostic)

			if tt.shouldSucceed {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}
		})
	}
}

func TestDiagnosticService_GetDiagnostics(t *testing.T) {
	logger := zap.NewNop()

	tests := []struct {
		name          string
		filter        *entities.DiagnosticFilter
		setupMocks    func(dr *mocks.MockDiagnosticRepository)
		expectedCount int
		expectedError string
		shouldSucceed bool
	}{
		{
			name:   "successful retrieval with no filters",
			filter: &entities.DiagnosticFilter{},
			setupMocks: func(dr *mocks.MockDiagnosticRepository) {
				dr.GetDiagnosticsFn = func(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error) {
					return []*entities.Diagnostic{
						{
							ID:        "1",
							PatientID: "0f6f6fce-127e-4e83-b865-e78bcbab881d",
							Diagnosis: "Common cold",
							Date:      time.Now(),
						},
						{
							ID:        "2",
							PatientID: "0f6f6fce-127e-4e83-b865-e78bcbab881d",
							Diagnosis: "Fever",
							Date:      time.Now(),
						},
					}, nil
				}
			},
			expectedCount: 2,
			shouldSucceed: true,
		},
		{
			name: "successful retrieval with patient name filter",
			filter: &entities.DiagnosticFilter{
				PatientName: stringPtr("John"),
			},
			setupMocks: func(dr *mocks.MockDiagnosticRepository) {
				dr.GetDiagnosticsFn = func(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error) {
					return []*entities.Diagnostic{
						{
							ID:        "1",
							PatientID: "0f6f6fce-127e-4e83-b865-e78bcbab881d",
							Diagnosis: "Common cold",
							Date:      time.Now(),
						},
					}, nil
				}
			},
			expectedCount: 1,
			shouldSucceed: true,
		},
		{
			name:   "database error",
			filter: &entities.DiagnosticFilter{},
			setupMocks: func(dr *mocks.MockDiagnosticRepository) {
				dr.GetDiagnosticsFn = func(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error) {
					return nil, fmt.Errorf("database connection error")
				}
			},
			expectedError: "database connection error",
			shouldSucceed: false,
		},
		{
			name: "empty result",
			filter: &entities.DiagnosticFilter{
				PatientName: stringPtr("NonExistent"),
			},
			setupMocks: func(dr *mocks.MockDiagnosticRepository) {
				dr.GetDiagnosticsFn = func(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error) {
					return []*entities.Diagnostic{}, nil
				}
			},
			expectedCount: 0,
			shouldSucceed: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDiagRepo := &mocks.MockDiagnosticRepository{}
			tt.setupMocks(mockDiagRepo)

			service := &DiagnosticService{
				diagnosticRepository: mockDiagRepo,
				logger:               logger,
			}

			diagnostics, err := service.GetDiagnostics(context.Background(), tt.filter)

			if tt.shouldSucceed {
				assert.NoError(t, err)
				assert.Len(t, diagnostics, tt.expectedCount)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, diagnostics)
			}
		})
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}
