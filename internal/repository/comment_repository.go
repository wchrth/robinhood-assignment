package repository

import (
	"robinhood-assignment/internal/model"

	"github.com/jmoiron/sqlx"
)

type CommentRepository interface {
	Create(comment *model.Comment) error
	Update(comment *model.Comment) error
	Delete(id int64) error
	GetByID(id int64) (*model.Comment, error)
	GetByAppointmentID(appointmentID int64) ([]model.Comment, error)
}

type commentRepositoryDB struct {
	db *sqlx.DB
}

func NewCommentRepositoryDB(db *sqlx.DB) CommentRepository {
	return &commentRepositoryDB{db: db}
}

func (repo *commentRepositoryDB) Create(comment *model.Comment) error {
	query := `
		INSERT INTO comments (description, user_id, appointment_id, created_date, created_by, updated_date, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err := repo.db.QueryRowx(query, comment.Description, comment.UserID, comment.AppointmentID, comment.CreatedDate, comment.CreatedBy, comment.UpdatedDate, comment.UpdatedBy).Scan(&comment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *commentRepositoryDB) Update(comment *model.Comment) error {
	query := `
		UPDATE comments SET description = $1, updated_date = $2, updated_by = $3
		WHERE id = $4`

	_, err := repo.db.Exec(query, comment.Description, comment.UpdatedDate, comment.UpdatedBy, comment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *commentRepositoryDB) Delete(id int64) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *commentRepositoryDB) GetByID(id int64) (*model.Comment, error) {
	var comment model.Comment
	query := "SELECT * FROM comments WHERE id=$1"
	err := repo.db.Get(&comment, query, id)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (repo *commentRepositoryDB) GetByAppointmentID(appointmentID int64) ([]model.Comment, error) {
	var comments []model.Comment
	query := `SELECT * FROM comments where appointment_id = $1`
	err := repo.db.Select(&comments, query, appointmentID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
