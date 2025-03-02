package entities

import "time"

type DiagnosticFilter struct {
	PatientID   *string    `json:"patientId"`
	PatientName *string    `json:"patientName"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	DiagnosisID *string    `json:"diagnosisId"`
	Status      *string    `json:"status"`
}

type Diagnostic struct {
	ID           string    `json:"id"`
	PatientID    string    `json:"patientId"`
	Date         time.Time `json:"date"`
	DiagnosisID  string    `json:"diagnosisId"`
	Status       string    `json:"status"`
	Prescription *string   `json:"prescription"`
}
