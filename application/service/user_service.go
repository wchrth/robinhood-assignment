package service

import (
	"robinhood-assignment/application/dto"
	"robinhood-assignment/domain/entity"
	"robinhood-assignment/domain/repository"
)

type UserService interface {
	GetByID(id int) (*dto.UserResponse, error)
	GetAll(offset, limit int) ([]dto.UserResponse, error)
	Create(request *dto.CreateUserRequest) error
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(ar repository.UserRepository) UserService {
	return userService{UserRepository: ar}
}

func (as userService) GetByID(id int) (*dto.UserResponse, error) {
	user, err := as.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	userResponse := dto.NewUserResponse(*user)

	return &userResponse, nil
}

func (as userService) GetAll(offset, limit int) ([]dto.UserResponse, error) {
	users, err := as.UserRepository.GetAll(offset, limit)
	if err != nil {
		return nil, err
	}

	userResponses := make([]dto.UserResponse, 0)
	for _, user := range users {
		userResponses = append(userResponses, dto.NewUserResponse(user))
	}

	return userResponses, nil
}

func (as userService) Create(request *dto.CreateUserRequest) error {

	user := &entity.User{
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	return as.UserRepository.Save(user)
}
