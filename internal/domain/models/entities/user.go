package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt string    `db:"created_at" json:"created_at"`
	UpdatedAt string    `db:"updated_at" json:"updated_at"`
	Credentials
}

type Credentials struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
