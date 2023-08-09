package postgres

import (
	"robinhood-assignment/domain/entity"
	"robinhood-assignment/domain/repository"

	"gorm.io/gorm"
)

type appointmentRepository struct {
	DB *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) repository.AppointmentRepository {
	return appointmentRepository{DB: db}
}

func (ar appointmentRepository) GetByID(id int) (*entity.Appointment, error) {
	appointment := &entity.Appointment{}
	if err := ar.DB.Preload("Comments").Preload("User").First(&appointment, id).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (ar appointmentRepository) GetAll(offset, limit int) ([]entity.Appointment, error) {
	appointments := []entity.Appointment{}
	if err := ar.DB.Preload("Comments").Preload("User").Offset(offset).Limit(limit).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (ar appointmentRepository) Save(appointment *entity.Appointment) error {
	if err := ar.DB.Save(&appointment).Error; err != nil {
		return err
	}
	return nil
}
