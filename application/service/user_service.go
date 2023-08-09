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

func NewUserService(ur repository.UserRepository) UserService {
	return userService{UserRepository: ur}
}

func (us userService) GetByID(id int) (*dto.UserResponse, error) {
	user, err := us.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	userResponse := dto.NewUserResponse(*user)

	return &userResponse, nil
}

func (us userService) GetAll(offset, limit int) ([]dto.UserResponse, error) {
	users, err := us.UserRepository.GetAll(offset, limit)
	if err != nil {
		return nil, err
	}

	userResponses := make([]dto.UserResponse, 0)
	for _, user := range users {
		userResponses = append(userResponses, dto.NewUserResponse(user))
	}

	return userResponses, nil
}

func (us userService) Create(request *dto.CreateUserRequest) error {

	user := &entity.User{
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	return us.UserRepository.Save(user)
}
