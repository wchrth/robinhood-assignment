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

func (ur userRepository) GetByID(id int) (*entity.User, error) {
	user := &entity.User{}
	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur userRepository) GetAll(offset, limit int) ([]entity.User, error) {
	users := []entity.User{}
	if err := ur.DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur userRepository) Save(user *entity.User) error {
	if err := ur.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
