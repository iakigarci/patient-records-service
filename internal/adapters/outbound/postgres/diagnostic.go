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

func (r *DiagnosticRepository) GetDiagnostics(ctx context.Context, filter *entities.DiagnosticFilter) ([]*entities.Diagnostic, error) {
	query := NewQueryBuilder().
		Query(BASE_DIAGNOSTIC_QUERY).
		Where("d.diagnosis_date >= $1").
		AddArgs(filter.StartDate).
		OrderBy("d.diagnosis_date DESC")

	var args []interface{}
	var conditions []string
	argPosition := 1

	if filter.PatientName != nil && *filter.PatientName != "" {
		conditions = append(conditions, fmt.Sprintf("p.name ILIKE $%d", argPosition))
		args = append(args, "%"+*filter.PatientName+"%")
		argPosition++
	}

	if filter.StartDate != nil {
		conditions = append(conditions, fmt.Sprintf("d.diagnosis_date >= $%d", argPosition))
		args = append(args, *filter.StartDate)
		argPosition++
	}

	if filter.EndDate != nil {
		conditions = append(conditions, fmt.Sprintf("d.diagnosis_date <= $%d", argPosition))
		args = append(args, *filter.EndDate)
		argPosition++
	}

	if len(conditions) > 0 {
		query.Where(strings.Join(conditions, " AND "))
	}

	rows, err := r.db.QueryxContext(ctx, query.Build())
	if err != nil {
		return nil, fmt.Errorf("error querying diagnoses: %v", err)
	}
	defer rows.Close()

	var diagnoses []*entities.Diagnostic
	for rows.Next() {
		var d entities.Diagnostic
		if err := rows.StructScan(&d); err != nil {
			return nil, fmt.Errorf("error scanning diagnosis row: %v", err)
		}
		diagnoses = append(diagnoses, &d)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating diagnosis rows: %v", err)
	}

	return diagnoses, nil
}
