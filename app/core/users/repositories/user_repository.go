package repositories

import (
	users "service-api/app/core/users/entities"
)

type UserRepository interface {
	FindAll() ([]users.User, error)
	Save(user users.User) (users.User, error)
	FindByEmail(email string) (users.User, error)
	FindByID(id uint) (users.User, error)
}
