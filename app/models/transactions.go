package models

import (
	"time"
)

type Transaction struct {
	ID          int       `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	Phonenumber string    `db:"phonenumber" json:"phonenumber" validate:"required,lte=255"`
	Summary     int       `db:"summary" json:"summary"`
}
