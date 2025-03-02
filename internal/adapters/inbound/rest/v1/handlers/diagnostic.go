package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
)

// @Summary Get diagnostics
// @Description Retrieves diagnostic records with optional filtering by patient name and date
// @Tags diagnostics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token"
// @Param patientName query string false "Patient name to filter by" example:"John Smith"
// @Param date query string false "Date to filter by (YYYY-MM-DD)" example:"2024-03-20"
// @Success 200 {array} DiagnosticResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/diagnostics [get]
func (h *DiagnosticHandler) GetDiagnostic(c *gin.Context) {
	// Check for Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}

	var filter DiagnosticFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diagnostics, err := h.diagnosticService.GetDiagnostics(c.Request.Context(), &entities.DiagnosticFilter{
		PatientName: filter.PatientName,
		StartDate:   filter.Date,
		EndDate:     filter.Date,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnostics)
}

type DiagnosticResponse struct {
	// Status of the diagnostic
	Status string `json:"status" example:"completed"`
	// Name of the patient
	PatientName string `json:"patientName" example:"John Smith"`
	// Date of the diagnostic
	Date time.Time `json:"date" example:"2024-03-20T15:04:05Z"`
	// Diagnosis details
	Diagnosis string `json:"diagnosis" example:"High blood glucose levels detected. Blood sugar at 180 mg/dL."`
	// Prescribed treatment
	Prescription string `json:"prescription" example:"Prescribed Metformin 500mg twice daily"`
}

type DiagnosticFilter struct {
	PatientName *string    `form:"patientName"`
	Date        *time.Time `form:"date" time_format:"2006-01-02"`
}

type DiagnosticHandler struct {
	diagnosticService ports.DiagnosticService
}

func NewDiagnosticHandler(ds ports.DiagnosticService) *DiagnosticHandler {
	return &DiagnosticHandler{
		diagnosticService: ds,
	}
}
