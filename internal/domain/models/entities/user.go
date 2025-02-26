package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	Phone     *string   `db:"phone" json:"phone"`
	Address   *Address  `db:"address" json:"address"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Address struct {
	Street string `db:"street" json:"street"`
	City   string `db:"city" json:"city"`
	State  string `db:"state" json:"state"`
	Zip    string `db:"zip" json:"zip"`
}
