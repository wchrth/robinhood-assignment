package repository

import "robinhood-assignment/domain/entity"

type UserRepository interface {
	GetByID(id int) (*entity.User, error)
	GetAll(offset, limit int) ([]entity.User, error)
	Save(user *entity.User) error
}
