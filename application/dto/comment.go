package dto

import (
	"robinhood-assignment/domain/entity"
	"time"
)

type CreateCommentRequest struct {
	Comment       string
	AppointmentID uint
	UserID        uint
}

type CommentResponse struct {
	ID              uint
	Comment         string
	UserID          uint
	UserDisplayName string
	UserEmail       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewCommentResponse(comment entity.Comment) CommentResponse {

	commentResponse := CommentResponse{
		ID:              comment.ID,
		Comment:         comment.Comment,
		UserID:          comment.UserID,
		UserDisplayName: comment.User.DisplayName,
		UserEmail:       comment.User.Email,
		CreatedAt:       comment.CreatedAt,
		UpdatedAt:       comment.UpdatedAt,
	}

	return commentResponse
}
