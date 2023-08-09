package service

import (
	"robinhood-assignment/application/dto"
	"robinhood-assignment/domain/entity"
	"robinhood-assignment/domain/repository"
)

type AppointmentService interface {
	GetByID(id int) (*dto.AppointmentResponse, error)
	GetAll(offset, limit int) ([]dto.AppointmentResponse, error)
	Create(request *dto.CreateAppointmentRequest) error
	Archive(id int) (*dto.AppointmentResponse, error)
}

type appointmentService struct {
	AppointmentRepository repository.AppointmentRepository
}

func NewAppointmentService(ar repository.AppointmentRepository) AppointmentService {
	return appointmentService{AppointmentRepository: ar}
}

func (as appointmentService) GetByID(id int) (*dto.AppointmentResponse, error) {
	appointment, err := as.AppointmentRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	appointmentResponse := dto.NewAppointmentResponse(*appointment)

	return &appointmentResponse, nil
}

func (as appointmentService) GetAll(offset, limit int) ([]dto.AppointmentResponse, error) {
	appointments, err := as.AppointmentRepository.GetAll(offset, limit)
	if err != nil {
		return nil, err
	}

	appointmentResponses := make([]dto.AppointmentResponse, 0)
	for _, appointment := range appointments {
		appointmentResponses = append(appointmentResponses, dto.NewAppointmentResponse(appointment))
	}

	return appointmentResponses, nil
}

func (as appointmentService) Create(request *dto.CreateAppointmentRequest) error {

	appointment := &entity.Appointment{
		Description: request.Description,
		UserID:      request.UserID,
	}

	return as.AppointmentRepository.Save(appointment)
}

func (as appointmentService) Archive(id int) (*dto.AppointmentResponse, error) {
	appointment, err := as.AppointmentRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	appointment.IsArchived = true
	as.AppointmentRepository.Save(appointment)

	appointmentResponse := dto.NewAppointmentResponse(*appointment)

	return &appointmentResponse, nil
}
