package handlers

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"github.com/iakigarci/go-ddd-microservice-template/internal/utils"
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

type CreateDiagnosticRequest struct {
	PatientID    string    `json:"patient_id" binding:"required" example:"77696893-740b-402c-989a-ff699b81853c"`
	Date         time.Time `json:"date" binding:"required" example:"2024-03-20T15:04:05Z"`
	Diagnosis    string    `json:"diagnosis" binding:"required" example:"Common cold"`
	Prescription *string   `json:"prescription,omitempty" example:"Rest and fluids"`
}

// CreateDiagnostic godoc
// @Summary Create a new diagnostic
// @Description Create a new diagnostic for a specific patient
// @Tags diagnostics
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param diagnostic body CreateDiagnosticRequest true "Diagnostic information"
// @Success 201 {object} entities.Diagnostic
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/diagnostics [post]
func (h *DiagnosticHandler) CreateDiagnostic(c *gin.Context) {
	var req CreateDiagnosticRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	diagnostic := &entities.Diagnostic{
		PatientID:    req.PatientID,
		Date:         req.Date,
		Diagnosis:    req.Diagnosis,
		Prescription: utils.If(req.Prescription != nil, sql.NullString{String: *req.Prescription, Valid: true}, sql.NullString{}),
	}

	if err := h.diagnosticService.CreateDiagnostic(c.Request.Context(), diagnostic); err != nil {
		if strings.Contains(err.Error(), "patient not found") {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create diagnostic"})
		return
	}

	c.JSON(http.StatusCreated, diagnostic)
}

type DiagnosticHandler struct {
	diagnosticService ports.DiagnosticService
}

func NewDiagnosticHandler(ds ports.DiagnosticService) *DiagnosticHandler {
	return &DiagnosticHandler{
		diagnosticService: ds,
	}
}
