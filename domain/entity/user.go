package entity

import "time"

type User struct {
	ID          uint
	DisplayName string    `gorm:"size:100;not null"`
	Email       string    `gorm:"size:100;not null;unique"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
