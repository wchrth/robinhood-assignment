package repository

import (
	"robinhood-assignment/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id int64) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
}

type userRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
	return &userRepositoryDB{db: db}
}

func (repo *userRepositoryDB) GetAll() ([]model.User, error) {
	var users []model.User
	query := "SELECT * FROM users"
	err := repo.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userRepositoryDB) GetByID(id int64) (*model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE id=$1"
	err := repo.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepositoryDB) GetByEmail(email string) (*model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE email=$1"
	err := repo.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepositoryDB) Create(user *model.User) error {
	query := `
		INSERT INTO users (email, password, display_name, created_date, updated_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := repo.db.QueryRowx(query, user.Email, user.Password, user.DisplayName, user.CreatedDate, user.UpdatedDate).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
