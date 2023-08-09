package postgres

import (
	"robinhood-assignment/domain/entity"
	"robinhood-assignment/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return userRepository{DB: db}
}

func (ar userRepository) GetByID(id int) (*entity.User, error) {
	user := &entity.User{}
	if err := ar.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ar userRepository) GetAll(offset, limit int) ([]entity.User, error) {
	users := []entity.User{}
	if err := ar.DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ar userRepository) Save(user *entity.User) error {
	if err := ar.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
