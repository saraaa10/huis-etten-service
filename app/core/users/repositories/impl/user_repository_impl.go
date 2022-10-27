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
	error := repo.db.Preload("UserType").Find(&users).Error
	if error != nil {
		return nil, error
	}

	return users, nil
}

func (repo *userRepository) Save(user entities.User) (entities.User, error) {
	error := repo.db.Create(&user).Error
	if error != nil {
		return entities.User{}, error
	}

	return user, nil
}

func (repo *userRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	error := repo.db.Preload("UserType").Where("email = ?", email).First(&user).Error
	if error != nil {
		return entities.User{}, error
	}

	return user, nil
}

func (repo *userRepository) FindByID(id uint) (entities.User, error) {
	var user entities.User
	error := repo.db.Preload("UserType").Where("id = ?", id).First(&user).Error
	if error != nil {
		return user, error
	}

	return user, nil
}
