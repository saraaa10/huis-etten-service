package impl

import (
	"service-api/app/core/users/entities"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	error := repo.db.Find(&users).Error
	if error != nil {
		return nil, error
	}

	return users, nil
}
