package repository

import "robinhood-assignment/domain/entity"

type CommentRepository interface {
	GetByID(id int) (*entity.Comment, error)
	GetAll(offset, limit int) ([]entity.Comment, error)
	Save(comment *entity.Comment) error
}
