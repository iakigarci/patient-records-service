package entities

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type User struct {
	ID           string         `db:"id" json:"id"`
	Email        string         `db:"email" json:"email"`
	PasswordHash string         `db:"password_hash" json:"password_hash"`
	Phone        sql.NullString `db:"phone" json:"phone,omitempty"`
	Address      *Address       `db:"address,type:jsonb" json:"address,omitempty"`
	CreatedAt    float64        `db:"created_at" json:"created_at"`
	UpdatedAt    float64        `db:"updated_at" json:"updated_at"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

func (a *Address) Scan(value interface{}) error {
	if value == nil {
		*a = Address{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", value)
	}

	return json.Unmarshal(bytes, &a)
}
