package model

import "time"

type User struct {
	ID          int64     `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	DisplayName string    `json:"display_name" db:"display_name"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}
