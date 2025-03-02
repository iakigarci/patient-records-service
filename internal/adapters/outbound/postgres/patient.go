package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/jmoiron/sqlx"
)

type PatientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) GetPatientByID(ctx context.Context, id string) (*entities.Patient, error) {
	var patient entities.Patient

	query := NewQueryBuilder().
		Query(BASE_PATIENT_QUERY).
		Where("id = $1").
		AddArgs(id)

	println(fmt.Sprintf("Query %v", query.Build()))
	println(fmt.Sprintf("Args %v", query.GetArgs()))

	err := r.db.GetContext(ctx, &patient, query.Build(), query.GetArgs()...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting patient: %v", err)
	}

	return &patient, nil
}
