package api

import "fmt"

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
	ErrEmailAlreadyInUse   = NewAPIError(409, "Email already in use")
	ErrUserNotFound        = NewAPIError(404, "User not found")
	ErrInternalServerError = NewAPIError(500, "Internal server error")
	ErrInvalidInput        = NewAPIError(400, "Invalid input")
	ErrInvalidUserID       = NewAPIError(400, "Invalid user ID")
)
