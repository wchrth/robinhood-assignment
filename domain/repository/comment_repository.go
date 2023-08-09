package repository

import "robinhood-assignment/domain/entity"

type CommentRepository interface {
	GetByID(id int) (*entity.Comment, error)
	GetAll(limit, offset int) ([]entity.Comment, error)
	Save(comment *entity.Comment) error
}
