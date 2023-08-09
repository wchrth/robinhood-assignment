package repository

import "robinhood-assignment/domain/entity"

type AppointmentRepository interface {
	GetByID(id int) (*entity.Appointment, error)
	GetAll(limit, offset int) ([]entity.Appointment, error)
	Save(appointment *entity.Appointment) error
}
