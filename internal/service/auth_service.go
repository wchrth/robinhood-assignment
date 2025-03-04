package service

import (
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/dto"
	"robinhood-assignment/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Authenticate(loginRequestDTO *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error)
	RefreshToken(refreshRequestDTO *dto.RefreshRequestDTO) (*dto.RefreshResponseDTO, error)
}

type authServiceImpl struct {
	userRepo   repository.UserRepository
	jwtService JWTService
}

func NewAuthServiceImpl(userRepo repository.UserRepository, jwtService JWTService) AuthService {
	return &authServiceImpl{userRepo: userRepo, jwtService: jwtService}
}

func (service *authServiceImpl) Authenticate(loginRequestDTO *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error) {
	user, err := service.userRepo.GetByEmail(loginRequestDTO.Email)
	if err != nil {
		return nil, api.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequestDTO.Password))
	if err != nil {
		return nil, api.ErrInvalidPassword
	}

	accessToken, err := service.jwtService.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, api.ErrInternalServerError
	}

	refreshToken, err := service.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, api.ErrInternalServerError
	}

	loginResponse := &dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return loginResponse, nil
}

func (service *authServiceImpl) RefreshToken(refreshRequestDTO *dto.RefreshRequestDTO) (*dto.RefreshResponseDTO, error) {
	claims, err := service.jwtService.ValidateToken(refreshRequestDTO.RefreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	newAccessToken, err := service.jwtService.GenerateAccessToken(claims.UserID)
	if err != nil {
		return nil, api.ErrInternalServerError
	}

	refreshResponse := &dto.RefreshResponseDTO{
		AccessToken: newAccessToken,
	}

	return refreshResponse, nil
}
