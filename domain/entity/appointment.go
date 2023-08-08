package entity

import "time"

type Appointment struct {
	ID          uint
	Description string
	Comments    []Comment
	UserID      uint
	User        User
	IsArchived  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
