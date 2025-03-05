package dto

import (
	"robinhood-assignment/internal/constant"
	"robinhood-assignment/internal/model"
	"time"
)

type CreateAppointmentDTO struct {
	Title       string `json:"title" binding:"required,max=255"`
	Description string `json:"description" binding:"required"`
}

type UpdateAppointmentDTO struct {
	Title       string                     `json:"title" binding:"required,max=255"`
	Description string                     `json:"description" binding:"required"`
	Status      constant.AppointmentStatus `json:"status" binding:"required,max=50"`
}

type AppointmentResponseDTO struct {
	ID          int64                      `json:"id"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	Status      constant.AppointmentStatus `json:"status"`
	CreatedDate time.Time                  `json:"created_date"`
	CreatedBy   string                     `json:"created_by"`
	UpdatedDate time.Time                  `json:"updated_date"`
	UpdatedBy   string                     `json:"updated_by"`
}

func ConvertCreateAppointmentDTOToAppointmentModel(dto *CreateAppointmentDTO) *model.Appointment {
	return &model.Appointment{
		Title:       dto.Title,
		Description: dto.Description,
	}
}

func ConvertUpdateAppointmentDTOToAppointmentModel(dto *UpdateAppointmentDTO) *model.Appointment {
	return &model.Appointment{
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
	}
}

func ConvertAppointmentModelToAppointmentResponseDTO(appointment *model.Appointment) *AppointmentResponseDTO {
	return &AppointmentResponseDTO{
		ID:          appointment.ID,
		Title:       appointment.Title,
		Description: appointment.Description,
		Status:      appointment.Status,
		CreatedDate: appointment.CreatedDate,
		CreatedBy:   appointment.CreatedBy,
		UpdatedDate: appointment.UpdatedDate,
		UpdatedBy:   appointment.UpdatedBy,
	}
}

func ConvertAppointmentModelsToAppointmentResponseDTOs(appointments []model.Appointment) []AppointmentResponseDTO {
	var appointmentResponseDTOs []AppointmentResponseDTO
	for _, appointment := range appointments {
		appointmentResponseDTO := ConvertAppointmentModelToAppointmentResponseDTO(&appointment)
		appointmentResponseDTOs = append(appointmentResponseDTOs, *appointmentResponseDTO)
	}

	return appointmentResponseDTOs
}
