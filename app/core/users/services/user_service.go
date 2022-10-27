package services

import (
	"service-api/app/core/users/entities"
	"service-api/app/util/users"
)

type UserService interface {
	FindAll() ([]entities.User, error)
	RegisterUser(user users.RegisterUserInput) (entities.User, error)
	LoginUser(user users.LoginUserInput) (entities.User, error)
	GetUserByID(id uint) (entities.User, error)
}
