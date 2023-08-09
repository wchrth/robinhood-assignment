package api

import (
	"net/http"
	"robinhood-assignment/application/dto"
	"robinhood-assignment/application/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	AppointmentService service.AppointmentService
}

func NewAppointmentHandler(as service.AppointmentService) AppointmentHandler {
	return AppointmentHandler{AppointmentService: as}
}

func (ah AppointmentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}

	appointment, err := ah.AppointmentService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := NewResponse(http.StatusOK, appointment)
	c.JSON(response.StatusCode, response)
}

func (ah AppointmentHandler) GetAll(c *gin.Context) {
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

	appointments, err := ah.AppointmentService.GetAll(offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := NewResponse(http.StatusOK, appointments)
	c.JSON(response.StatusCode, response)
}

func (ah AppointmentHandler) Create(c *gin.Context) {
	request := dto.CreateAppointmentRequest{}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := ah.AppointmentService.Create(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := NewResponse(http.StatusCreated, request)
	c.JSON(response.StatusCode, response)

}
