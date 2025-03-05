package handler

import (
	"net/http"
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/constant"
	"robinhood-assignment/internal/dto"
	"robinhood-assignment/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	appointmentService service.AppointmentService
}

func NewAppointmentHandler(appointmentService service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{appointmentService}
}

func (h *AppointmentHandler) GetStatuses(c *gin.Context) {
	api.RespondSuccess(c, http.StatusOK, "Appointment statuses retrieved successfully", constant.AppointmentStatuses())
}

func (h *AppointmentHandler) GetAllAppointments(c *gin.Context) {
	appointments, err := h.appointmentService.GetAll()
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccess(c, http.StatusOK, "Appointments retrieved successfully", appointments)
}

func (h *AppointmentHandler) GetAppointmentByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	appointment, err := h.appointmentService.GetByID(id)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccess(c, http.StatusOK, "Appointment retrieved successfully", appointment)
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var createAppointmentDTO dto.CreateAppointmentDTO
	if err := c.ShouldBindJSON(&createAppointmentDTO); err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	userID := c.GetInt64("userID")

	err := h.appointmentService.Create(&createAppointmentDTO, userID)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccessNoData(c, http.StatusCreated, "Appointment created successfully")
}

func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	var updateAppointmentDTO dto.UpdateAppointmentDTO
	if err := c.ShouldBindJSON(&updateAppointmentDTO); err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	userID := c.GetInt64("userID")

	err = h.appointmentService.Update(&updateAppointmentDTO, id, userID)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccessNoData(c, http.StatusOK, "Appointment updated successfully")
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		apiErr := api.ErrInvalidInput
		api.RespondError(c, apiErr)
		return
	}

	err = h.appointmentService.Delete(id)
	if err != nil {
		api.RespondError(c, err)
		return
	}

	api.RespondSuccessNoData(c, http.StatusOK, "Appointment deleted successfully")
}
