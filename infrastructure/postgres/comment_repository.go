package postgres

import (
	"robinhood-assignment/domain/entity"
	"robinhood-assignment/domain/repository"

	"gorm.io/gorm"
)

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return commentRepository{DB: db}
}

func (ar commentRepository) GetByID(id int) (*entity.Comment, error) {
	comment := &entity.Comment{}
	if err := ar.DB.Preload("Appointment").Preload("User").First(&comment, id).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (ar commentRepository) GetAll(offset, limit int) ([]entity.Comment, error) {
	comments := []entity.Comment{}
	if err := ar.DB.Preload("Appointment").Preload("User").Offset(offset).Limit(limit).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (ar commentRepository) Save(comment *entity.Comment) error {
	if err := ar.DB.Save(&comment).Error; err != nil {
		return err
	}
	return nil
}
