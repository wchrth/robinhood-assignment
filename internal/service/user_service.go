package service

import (
	"database/sql"
	"errors"
	"fmt"
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/dto"
	"robinhood-assignment/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll() ([]dto.UserResponseDTO, error)
	GetByID(id int64) (*dto.UserResponseDTO, error)
	GetByEmail(email string) (*dto.UserResponseDTO, error)
	Create(userDTO *dto.UserDTO) error
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserServiceImpl(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (service *userServiceImpl) GetAll() ([]dto.UserResponseDTO, error) {
	users, err := service.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return dto.ToUserResponseDTOs(users), nil
}

func (service *userServiceImpl) GetByID(id int64) (*dto.UserResponseDTO, error) {
	user, err := service.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api.ErrUserNotFound
		}
		return nil, err
	}

	return dto.ToUserResponseDTO(user), nil
}

func (service *userServiceImpl) GetByEmail(email string) (*dto.UserResponseDTO, error) {
	user, err := service.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api.ErrUserNotFound
		}
		return nil, err
	}

	return dto.ToUserResponseDTO(user), nil
}

func (service *userServiceImpl) Create(userDTO *dto.UserDTO) error {

	user := dto.ToUserModel(userDTO)

	// Check if the user already exists by email
	existingUser, err := service.userRepo.GetByEmail(user.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	if err == nil && existingUser != nil {
		return api.ErrEmailAlreadyInUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.CreatedDate = time.Now().UTC()
	user.UpdatedDate = time.Now().UTC()

	err = service.userRepo.Create(user)
	if err != nil {
		return err
	}
	fmt.Println(user)

	return nil
}
