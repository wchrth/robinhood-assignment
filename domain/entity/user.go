package entity

import "time"

type User struct {
	ID          uint
	DisplayName string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
