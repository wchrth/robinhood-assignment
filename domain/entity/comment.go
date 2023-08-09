package entity

import "time"

type Comment struct {
	ID            uint
	Comment       string `gorm:"size:200;not null"`
	AppointmentID uint
	Appointment   Appointment
	UserID        uint
	User          User
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
