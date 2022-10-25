package services

import "service-api/app/core/menu/entities"

type MenuService interface {
	FindAll() ([]entities.Menu, error)
}
