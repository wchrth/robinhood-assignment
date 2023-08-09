package dto

import (
	"robinhood-assignment/domain/entity"
	"time"
)

type CreateCommentRequest struct {
	Comment       string `json:"comment"`
	AppointmentID uint   `json:"appointmentID"`
	UserID        uint   `json:"userID"`
}

type CommentResponse struct {
	ID              uint      `json:"id"`
	Comment         string    `json:"comment"`
	UserID          uint      `json:"userID"`
	UserDisplayName string    `json:"userDisplayName"`
	UserEmail       string    `json:"userEmail"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
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
