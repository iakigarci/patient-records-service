package postgres

import (
	"strings"
)

var (
	BASE_USER_QUERY = `
		SELECT id, 
			email, 
			password_hash,
			address,
			phone,
			EXTRACT(EPOCH FROM created_at) as created_at,
			EXTRACT(EPOCH FROM updated_at) as updated_at
		FROM users`

	BASE_DIAGNOSTIC_QUERY = `
		SELECT id, 
			diagnosis,
			prescription,
			diagnosis_date,
			created_at,
			updated_at,
			patient_id,
			patient_name,
			patient_dni,
			patient_email,
			patient_phone,
			patient_address
		FROM diagnoses`
)

type QueryBuilder struct {
	query      string
	where      []string
	orderBy    string
	pagination string
	args       []any
	paramCount int
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		where: make([]string, 0, 4),
		args:  make([]any, 0, 4),
	}
}

func (qb *QueryBuilder) Query(query string) *QueryBuilder {
	qb.query = query
	return qb
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
	qb.where = append(qb.where, condition)
	return qb
}

func (qb *QueryBuilder) OrderBy(orderBy string) *QueryBuilder {
	qb.orderBy = orderBy
	return qb
}

func (qb *QueryBuilder) Paginate(pagination string) *QueryBuilder {
	qb.pagination = pagination
	return qb
}

func (qb *QueryBuilder) AddArgs(args ...any) *QueryBuilder {
	for _, arg := range args {
		qb.paramCount++
		qb.args = append(qb.args, arg)
	}
	return qb
}

func (qb *QueryBuilder) GetArgs() []any {
	return qb.args
}

func (qb *QueryBuilder) Build() string {
	var query strings.Builder
	query.WriteString(qb.query)

	if len(qb.where) > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(strings.Join(qb.where, " AND "))
	}

	if qb.orderBy != "" {
		query.WriteString(" ORDER BY ")
		query.WriteString(qb.orderBy)
	}

	if qb.pagination != "" {
		query.WriteString(" ")
		query.WriteString(qb.pagination)
	}

	return query.String()
}
