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
	ID           string         `db:"id" json:"id"`
	PatientID    string         `db:"patient_id" json:"patient_id"`
	Date         float64        `db:"diagnosis_date" json:"diagnosis_date"`
	Diagnosis    string         `db:"diagnosis" json:"diagnosis"`
	Prescription sql.NullString `db:"prescription" json:"prescription"`
	CreatedAt    float64        `db:"created_at" json:"created_at"`
	UpdatedAt    float64        `db:"updated_at" json:"updated_at"`
}
