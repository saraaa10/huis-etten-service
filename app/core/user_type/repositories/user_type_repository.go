package repositories

import (
	"service-api/app/core/user_type/entities"
)

type UserTypeRepository interface {
	FindAll() ([]entities.UserType, error)
}
