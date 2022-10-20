package repositories

import "service-api/app/core/type_menu/entities"

type TypeMenuRepository interface {
	FindAll() ([]entities.TypeMenu, error)
}
