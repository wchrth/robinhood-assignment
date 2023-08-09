package dto

import (
	"robinhood-assignment/domain/entity"
	"time"
)

type CreateUserRequest struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

type UserResponse struct {
	ID          uint      `json:"id"`
	DisplayName string    `json:"displayName"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewUserResponse(user entity.User) UserResponse {

	userResponse := UserResponse{
		ID:          user.ID,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return userResponse
}
