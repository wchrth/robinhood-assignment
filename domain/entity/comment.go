package entity

import "time"

type Comment struct {
	ID            uint
	Comment       string
	AppointmentID uint
	Appointment   Appointment
	UserID        uint
	User          User
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
