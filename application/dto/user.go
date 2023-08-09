package dto

import (
	"robinhood-assignment/domain/entity"
	"time"
)

type CreateUserRequest struct {
	DisplayName string
	Email       string
}

type UserResponse struct {
	ID          uint
	DisplayName string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
