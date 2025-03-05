package service

import (
	"database/sql"
	"errors"
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/constant"
	"robinhood-assignment/internal/dto"
	"robinhood-assignment/internal/model"
	"robinhood-assignment/internal/repository"
	"time"
)

type AppointmentService interface {
	GetAll() ([]dto.AppointmentResponseDTO, error)
	GetByID(id int64) (*dto.AppointmentResponseDTO, error)
	Create(createAppointmentDTO *dto.CreateAppointmentDTO, userID int64) error
	Update(updateAppointmentDTO *dto.UpdateAppointmentDTO, id, userID int64) error
	Delete(id int64) error
	Archive(id, userID int64) error
	GetHistories(appointmentID int64) ([]model.AppointmentHistory, error)
}

type appointmentServiceImpl struct {
	appointmentRepo repository.AppointmentRepository
	userRepo        repository.UserRepository
}

func NewAppointmentServiceImpl(appointmentRepo repository.AppointmentRepository, userRepo repository.UserRepository) AppointmentService {
	return &appointmentServiceImpl{appointmentRepo, userRepo}
}

func (service *appointmentServiceImpl) GetAll() ([]dto.AppointmentResponseDTO, error) {
	appointments, err := service.appointmentRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return dto.ConvertAppointmentModelsToAppointmentResponseDTOs(appointments), nil
}

func (service *appointmentServiceImpl) GetByID(id int64) (*dto.AppointmentResponseDTO, error) {
	appointment, err := service.appointmentRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api.ErrDataNotFound
		}
		return nil, err
	}

	return dto.ConvertAppointmentModelToAppointmentResponseDTO(appointment), nil
}

func (service *appointmentServiceImpl) Create(createAppointmentDTO *dto.CreateAppointmentDTO, userID int64) error {

	user, err := service.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrUserNotFound
		}
		return err
	}

	appointment := dto.ConvertCreateAppointmentDTOToAppointmentModel(createAppointmentDTO)

	appointment.Status = constant.ToDo
	appointment.IsArchived = false
	appointment.UserID = userID
	appointment.CreatedDate = time.Now().UTC()
	appointment.CreatedBy = user.DisplayName
	appointment.UpdatedDate = time.Now().UTC()
	appointment.UpdatedBy = user.DisplayName

	err = service.appointmentRepo.Create(appointment)
	if err != nil {
		return err
	}

	var history model.AppointmentHistory
	history.Title = appointment.Title
	history.Description = appointment.Description
	history.Status = appointment.Status
	history.AppointmentID = appointment.ID

	err = service.appointmentRepo.CreateHistory(&history)
	if err != nil {
		return err
	}

	return nil
}

func (service *appointmentServiceImpl) Update(updateAppointmentDTO *dto.UpdateAppointmentDTO, id, userID int64) error {

	if !updateAppointmentDTO.Status.IsValid() {
		return api.ErrInvalidInput
	}

	appointment, err := service.appointmentRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrDataNotFound
		}
		return err
	}

	user, err := service.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrUserNotFound
		}
		return err
	}

	appointment.Title = updateAppointmentDTO.Title
	appointment.Description = updateAppointmentDTO.Description
	appointment.Status = updateAppointmentDTO.Status
	appointment.UpdatedDate = time.Now().UTC()
	appointment.UpdatedBy = user.DisplayName

	err = service.appointmentRepo.Update(appointment)
	if err != nil {
		return err
	}

	var history model.AppointmentHistory
	history.Title = appointment.Title
	history.Description = appointment.Description
	history.Status = appointment.Status
	history.AppointmentID = appointment.ID

	err = service.appointmentRepo.CreateHistory(&history)
	if err != nil {
		return err
	}

	return nil
}

func (service *appointmentServiceImpl) Delete(id int64) error {
	_, err := service.appointmentRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrDataNotFound
		}
		return err
	}

	err = service.appointmentRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (service *appointmentServiceImpl) Archive(id, userID int64) error {
	appointment, err := service.appointmentRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrDataNotFound
		}
		return err
	}

	user, err := service.userRepo.GetByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.ErrUserNotFound
		}
		return err
	}

	appointment.IsArchived = true
	appointment.UpdatedDate = time.Now().UTC()
	appointment.UpdatedBy = user.DisplayName

	err = service.appointmentRepo.Update(appointment)
	if err != nil {
		return err
	}

	return nil
}

func (service *appointmentServiceImpl) GetHistories(appointmentID int64) ([]model.AppointmentHistory, error) {
	histories, err := service.appointmentRepo.GetHistories(appointmentID)
	if err != nil {
		return nil, err
	}
	return histories, nil
}
