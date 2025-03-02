package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
)

type DiagnosticResponse struct {
	Status      string    `json:"status" example:"ok"`
	PatientName string    `json:"patientName" example:"John Doe"`
	Date        time.Time `json:"date" example:"2024-03-20T15:04:05Z"`
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

func (h *DiagnosticHandler) GetDiagnostic(c *gin.Context) {
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
