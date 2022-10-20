package services

import "service-api/app/core/type_menu/entities"

type TypeMenuService interface {
	FindAll() ([]entities.TypeMenu, error)
}
