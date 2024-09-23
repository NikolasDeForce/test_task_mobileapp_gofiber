package models

import (
	"time"
)

type User struct {
	ID          int       `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	FName       string    `db:"fname" json:"fname" validate:"required,lte=255"`
	Email       string    `db:"email" json:"email" validate:"required,lte=255"`
	Phonenumber string    `db:"phonenumber" json:"phonenumber" validate:"required,lte=255"`
	Password    string    `db:"password" json:"password" validate:"required,lte=255"`
	Gender      string    `db:"gender" json:"gender" validate:"required,lte=255"`
	Birthday    string    `db:"birthday" json:"birthday" validate:"required,lte=255"`
	Balance     int       `db:"balance" json:"balance"`
}
