package model

import (
	"robinhood-assignment/internal/constant"
	"time"
)

type Appointment struct {
	ID          int64                      `json:"id" db:"id"`
	Title       string                     `json:"title" db:"title"`
	Description string                     `json:"description" db:"description"`
	Status      constant.AppointmentStatus `json:"status" db:"status"`
	IsArchived  bool                       `json:"is_archived" db:"is_archived"`
	UserID      int64                      `json:"user_id" db:"user_id"`
	CreatedDate time.Time                  `json:"created_date" db:"created_date"`
	CreatedBy   string                     `json:"created_by" db:"created_by"`
	UpdatedDate time.Time                  `json:"updated_date" db:"updated_date"`
	UpdatedBy   string                     `json:"updated_by" db:"updated_by"`
}

type AppointmentHistory struct {
	ID            int64                      `json:"id" db:"id"`
	Title         string                     `json:"title" db:"title"`
	Description   string                     `json:"description" db:"description"`
	Status        constant.AppointmentStatus `json:"status" db:"status"`
	AppointmentID int64                      `json:"appointment_id" db:"appointment_id"`
}
