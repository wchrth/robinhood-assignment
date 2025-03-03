package handler

import (
	"net/http"
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/dto"
	"robinhood-assignment/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAll()
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccess(c, http.StatusOK, "Users retrieved successfully", users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		apiErr := api.ErrInvalidUserID
		api.RespondError(c, apiErr)
		return
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccess(c, http.StatusOK, "User retrieved successfully", user)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {

	email := c.Param("email")

	user, err := h.userService.GetByEmail(email)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccess(c, http.StatusOK, "User retrieved successfully", user)
}

func (h *UserHandler) Register(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	err := h.userService.Create(&userDTO)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccessNoData(c, http.StatusCreated, "User created successfully")
}
