package repositories

import "service-api/app/core/menu/entities"

type MenuRepository interface {
	FindAll() ([]entities.Menu, error)
}