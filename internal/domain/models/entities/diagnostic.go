package entities

import (
	"database/sql"
	"time"
)

type DiagnosticFilter struct {
	PatientID   *string    `json:"patientId"`
	PatientName *string    `json:"patientName"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	DiagnosisID *string    `json:"diagnosisId"`
	Status      *string    `json:"status"`
}

type Diagnostic struct {
	ID           string         `json:"id" db:"id"`
	PatientID    string         `json:"patient_id" db:"patient_id"`
	Date         time.Time      `json:"date" db:"diagnosis_date"`
	Diagnosis    string         `json:"diagnosis" db:"diagnosis"`
	Prescription sql.NullString `json:"prescription,omitempty" db:"prescription"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at"`
}
