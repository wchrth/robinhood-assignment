package postgres

import (
	"robinhood-assignment/domain/entity"
	"robinhood-assignment/domain/repository"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return CommentRepository{DB: db}
}

func (ar CommentRepository) GetByID(id int) (*entity.Comment, error) {
	comment := &entity.Comment{}
	if err := ar.DB.Preload("Appointment").First(&comment, id).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (ar CommentRepository) GetAll(offset, limit int) ([]entity.Comment, error) {
	comments := []entity.Comment{}
	if err := ar.DB.Preload("Appointment").Offset(offset).Limit(limit).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (ar CommentRepository) Save(comment *entity.Comment) error {
	if err := ar.DB.Save(&comment).Error; err != nil {
		return err
	}
	return nil
}
