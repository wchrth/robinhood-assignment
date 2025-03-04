package api

import (
	"fmt"
	"net/http"
)

type APIError struct {
	Code    int    // HTTP status code (e.g., 404, 500)
	Message string // Error message
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

var (
	ErrEmailAlreadyInUse   = NewAPIError(http.StatusConflict, "Email already in use")
	ErrUserNotFound        = NewAPIError(http.StatusNotFound, "User not found")
	ErrInternalServerError = NewAPIError(http.StatusInternalServerError, "Internal server error")
	ErrInvalidInput        = NewAPIError(http.StatusBadRequest, "Invalid input")
	ErrInvalidUserID       = NewAPIError(http.StatusBadRequest, "Invalid user ID")
	ErrInvalidToken        = NewAPIError(http.StatusUnauthorized, "Invalid token")
	ErrInvalidClaims       = NewAPIError(http.StatusUnauthorized, "Invalid claims")
	ErrInvalidTokenType    = NewAPIError(http.StatusUnauthorized, "Invalid token type")
	ErrInvalidPassword     = NewAPIError(http.StatusUnauthorized, "Invalid password")
	ErrUnauthorized        = NewAPIError(http.StatusUnauthorized, "Unauthorized")
)
