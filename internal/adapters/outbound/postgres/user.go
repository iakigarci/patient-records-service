package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iakigarci/go-ddd-microservice-template/internal/domain/models/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (db *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := NewQueryBuilder().
		Query(BASE_USER_QUERY).
		Where("email = $1").
		AddArgs(email)

	row := db.db.QueryRowContext(ctx, query.Build(), query.GetArgs()...)

	var user entities.User
	err := row.Scan(&user)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
