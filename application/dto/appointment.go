package dto

import (
	"robinhood-assignment/domain/entity"
	"time"
)

type CreateAppointmentRequest struct {
	Description string `json:"description"`
	UserID      uint   `json:"userID"`
}

type AppointmentResponse struct {
	ID              uint              `json:"id"`
	Description     string            `json:"description"`
	Comments        []CommentResponse `json:"comments"`
	UserID          uint              `json:"userID"`
	UserDisplayName string            `json:"userDisplayName"`
	UserEmail       string            `json:"userEmail"`
	IsArchived      bool              `json:"isArchived"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
}

func NewAppointmentResponse(appointment entity.Appointment) AppointmentResponse {
	comments := make([]CommentResponse, 0)
	for _, comment := range appointment.Comments {
		c := CommentResponse{
			ID:              comment.ID,
			Comment:         comment.Comment,
			UserID:          comment.UserID,
			UserDisplayName: comment.User.DisplayName,
			UserEmail:       comment.User.Email,
			CreatedAt:       comment.CreatedAt,
			UpdatedAt:       comment.UpdatedAt,
		}

		comments = append(comments, c)
	}

	appointmentResponse := AppointmentResponse{
		ID:              appointment.ID,
		Description:     appointment.Description,
		Comments:        comments,
		UserID:          appointment.UserID,
		UserDisplayName: appointment.User.DisplayName,
		UserEmail:       appointment.User.Email,
		IsArchived:      appointment.IsArchived,
		CreatedAt:       appointment.CreatedAt,
		UpdatedAt:       appointment.UpdatedAt,
	}

	return appointmentResponse
}
