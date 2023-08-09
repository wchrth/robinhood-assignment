package dto

import "time"

type CommentResponse struct {
	ID              uint
	Comment         string
	UserID          uint
	UserDisplayName string
	UserEmail       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
