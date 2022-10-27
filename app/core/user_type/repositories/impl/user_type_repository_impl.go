package impl

import (
	"service-api/app/core/user_type/entities"

	"gorm.io/gorm"
)

type userTypeRepository struct {
	db *gorm.DB
}

func NewUserTypeRepository(db *gorm.DB) *userTypeRepository {
	return &userTypeRepository{
		db: db,
	}
}

func (repo *userTypeRepository) FindAll() ([]entities.UserType, error) {
	var userTypes []entities.UserType
	err := repo.db.Find(&userTypes).Error

	if err != nil {
		return nil, err
	}

	return userTypes, nil
}

func (repo *userTypeRepository) FindUserTypeById(id uint) (entities.UserType, error) {
	var userType entities.UserType
	err := repo.db.Where("id = ?", id).Find(&userType).Error
	if err != nil {
		return userType, err
	}

	return userType, nil
}