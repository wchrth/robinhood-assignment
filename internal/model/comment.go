package model

import "time"

type Comment struct {
	ID            int64     `json:"id" db:"id"`
	Description   string    `json:"description" db:"description"`
	UserID        int64     `json:"user_id" db:"user_id"`
	AppointmentID int64     `json:"appointment_id" db:"appointment_id"`
	CreatedDate   time.Time `json:"created_date" db:"created_date"`
	CreatedBy     string    `json:"created_by" db:"created_by"`
	UpdatedDate   time.Time `json:"updated_date" db:"updated_date"`
	UpdatedBy     string    `json:"updated_by" db:"updated_by"`
}
