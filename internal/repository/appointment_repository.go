package repository

import (
	"robinhood-assignment/internal/model"

	"github.com/jmoiron/sqlx"
)

type AppointmentRepository interface {
	GetAll() ([]model.Appointment, error)
	GetByID(id int64) (*model.Appointment, error)
	Create(appointment *model.Appointment) error
	Update(appointment *model.Appointment) error
	Delete(id int64) error
}

type appointmentRepositoryDB struct {
	db *sqlx.DB
}

func NewAppointmentRepositoryDB(db *sqlx.DB) AppointmentRepository {
	return &appointmentRepositoryDB{db: db}
}

func (repo *appointmentRepositoryDB) GetAll() ([]model.Appointment, error) {
	var appointments []model.Appointment
	query := `SELECT * FROM appointments where is_archived = false`
	err := repo.db.Select(&appointments, query)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (repo *appointmentRepositoryDB) GetByID(id int64) (*model.Appointment, error) {
	var appointment model.Appointment
	query := "SELECT * FROM appointment WHERE id=$1"
	err := repo.db.Get(&appointment, query, id)
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (repo *appointmentRepositoryDB) Create(appointment *model.Appointment) error {
	query := `
		INSERT INTO appointments (title, description, status, is_archived, user_id, created_date, created_by, updated_date, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`

	err := repo.db.QueryRowx(query, appointment.Title, appointment.Description, appointment.Status, appointment.IsArchived, appointment.UserID, appointment.CreatedDate, appointment.CreatedBy, appointment.UpdatedDate, appointment.UpdatedBy).Scan(&appointment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *appointmentRepositoryDB) Update(appointment *model.Appointment) error {
	query := `
		UPDATE appointments SET title = $1, description = $2, status = $3, is_archived = $4, 
		user_id = $5, updated_date = $6, updated_by = $7
		WHERE id = $8`

	_, err := repo.db.Exec(query, appointment.Title, appointment.Description, appointment.Status, appointment.IsArchived, appointment.UserID, appointment.UpdatedDate, appointment.UpdatedBy, appointment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *appointmentRepositoryDB) Delete(id int64) error {
	query := `DELETE FROM appointments WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
