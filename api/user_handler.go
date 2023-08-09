package api

import (
	"net/http"
	"robinhood-assignment/application/dto"
	"robinhood-assignment/application/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(as service.UserService) UserHandler {
	return UserHandler{UserService: as}
}

func (uh UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}

	appointment, err := uh.UserService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := NewResponse(http.StatusOK, appointment)
	c.JSON(response.StatusCode, response)
}

func (uh UserHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	appointments, err := uh.UserService.GetAll(offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := NewResponse(http.StatusOK, appointments)
	c.JSON(response.StatusCode, response)
}

func (uh UserHandler) Create(c *gin.Context) {
	request := dto.CreateUserRequest{}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := uh.UserService.Create(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := NewResponse(http.StatusCreated, request)
	c.JSON(response.StatusCode, response)

}
