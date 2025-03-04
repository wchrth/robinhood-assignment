package service

import (
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID    int64  `json:"user_id"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateAccessToken(userID int64) (string, error)
	GenerateRefreshToken(userID int64) (string, error)
	ValidateToken(tokenStr string, tokenType string) (*Claims, error)
}

type jwtServiceImpl struct {
	jwtConfig *config.JWTConfig
}

func NewJWTServiceImpl(jwtConfig *config.JWTConfig) JWTService {
	return &jwtServiceImpl{
		jwtConfig: jwtConfig,
	}
}

func (service *jwtServiceImpl) GenerateAccessToken(userID int64) (string, error) {
	claims := &Claims{
		UserID:    userID,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(service.jwtConfig.AccessExpirationTime)),
			Issuer:    service.jwtConfig.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(service.jwtConfig.SecretKey))
}

func (service *jwtServiceImpl) GenerateRefreshToken(userID int64) (string, error) {
	claims := &Claims{
		UserID:    userID,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(service.jwtConfig.RefreshExpirationTime)),
			Issuer:    service.jwtConfig.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(service.jwtConfig.SecretKey))
}

func (service *jwtServiceImpl) ValidateToken(tokenStr string, tokenType string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.jwtConfig.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, api.ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, api.ErrInvalidClaims
	}

	if claims.TokenType != tokenType {
		return nil, api.ErrInvalidTokenType
	}

	return claims, nil
}
