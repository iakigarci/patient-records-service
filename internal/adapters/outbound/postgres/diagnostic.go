package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/ports"
	"github.com/jmoiron/sqlx"
)

type DiagnosticRepository struct {
	db *sqlx.DB
}

func NewDiagnosticRepository(db *sqlx.DB) ports.DiagnosticRepository {
	return &DiagnosticRepository{
		db: db,
	}
}

func (db *DiagnosticRepository) GetDiagnostics(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error) {
	query := NewQueryBuilder().
		Query(BASE_DIAGNOSTIC_QUERY).
		OrderBy("d.diagnosis_date DESC")

	var args []interface{}
	var conditions []string
	argPosition := 1

	println(fmt.Sprintf("Filter %v", filter))

	if filter.PatientName != nil && *filter.PatientName != "" {
		println(fmt.Sprintf("Patient Name %v", *filter.PatientName))
		conditions = append(conditions, fmt.Sprintf("POSITION(LOWER($%d) IN LOWER(p.name)) > 0", argPosition))
		args = append(args, strings.ToLower(*filter.PatientName))
		argPosition++
	}

	if filter.StartDate != nil {
		conditions = append(conditions, fmt.Sprintf("DATE(diagnosis_date) >= DATE($%d::timestamp)", argPosition))
		args = append(args, filter.StartDate)
		argPosition++
	}

	if filter.EndDate != nil {
		conditions = append(conditions, fmt.Sprintf("DATE(diagnosis_date) <= DATE($%d::timestamp)", argPosition))
		args = append(args, filter.EndDate)
		argPosition++
	}

	if len(conditions) > 0 {
		query.Where(strings.Join(conditions, " AND "))
		query.AddArgs(args...)
	}

	println(fmt.Sprintf("Query %v", query.Build()))
	println(fmt.Sprintf("Args %v", query.GetArgs()))

	diagnostics, err := MultipleQuery[entities.Diagnostic](ctx, db.db, query.Build(), query.GetArgs()...)
	if err != nil {
		return nil, fmt.Errorf("error querying diagnoses: %v", err)
	}

	return diagnostics, nil
}
