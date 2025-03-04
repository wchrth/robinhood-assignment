package handler

import (
	"net/http"
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/dto"
	"robinhood-assignment/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequestDTO dto.LoginRequestDTO
	if err := c.ShouldBindJSON(&loginRequestDTO); err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	loginResponse, err := h.authService.Authenticate(&loginRequestDTO)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccess(c, http.StatusOK, "Login successfully", loginResponse)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var refreshRequestDTO dto.RefreshRequestDTO
	if err := c.ShouldBindJSON(&refreshRequestDTO); err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	refreshResponse, err := h.authService.RefreshToken(&refreshRequestDTO)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccess(c, http.StatusOK, "Refresh token successfully", refreshResponse)
}
