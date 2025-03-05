package dto

import (
	"robinhood-assignment/internal/model"
	"time"
)

type UserDTO struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	DisplayName string `json:"display_name" binding:"required,min=3,max=100"`
}

type UserResponseDTO struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

func ConvertUserDTOToUserModel(dto *UserDTO) *model.User {
	return &model.User{
		Email:       dto.Email,
		Password:    dto.Password,
		DisplayName: dto.DisplayName,
	}
}

func ConvertUserModelToUserResponseDTO(user *model.User) *UserResponseDTO {
	return &UserResponseDTO{
		ID:          user.ID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		CreatedDate: user.CreatedDate,
		UpdatedDate: user.UpdatedDate,
	}
}

func ConvertUserModelsToUserResponseDTOs(users []model.User) []UserResponseDTO {
	var userDTOs []UserResponseDTO
	for _, user := range users {
		userDTO := ConvertUserModelToUserResponseDTO(&user)
		userDTOs = append(userDTOs, *userDTO)
	}

	return userDTOs
}
