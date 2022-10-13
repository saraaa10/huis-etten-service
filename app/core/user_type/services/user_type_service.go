package services

import "service-api/app/core/user_type/entities"

type UserTypeService interface {
	GetAll() ([]entities.UserType, error)
}
