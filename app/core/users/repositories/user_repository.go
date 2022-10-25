package repositories

import "service-api/app/core/users/entities"

type UserRepository interface {
	FindAll() ([]entities.User, error)
}
