package entities

type Patient struct {
	ID        string  `db:"id" json:"id"`
	Name      string  `db:"name" json:"name"`
	DNI       string  `db:"dni" json:"dni"`
	Email     string  `db:"email" json:"email"`
	Phone     string  `db:"phone" json:"phone"`
	Address   string  `db:"address" json:"address"`
	CreatedAt float64 `db:"created_at" json:"created_at"`
	UpdatedAt float64 `db:"updated_at" json:"updated_at"`
}
