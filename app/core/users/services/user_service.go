package services

import "service-api/app/core/users/entities"

type UserService interface {
	FindAll() ([]entities.User)
}