package entity

import "time"

type Appointment struct {
	ID          uint
	Description string `gorm:"size:200;not null"`
	Comments    []Comment
	UserID      uint `gorm:"not null"`
	User        User
	IsArchived  bool      `gorm:"default:false;not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
