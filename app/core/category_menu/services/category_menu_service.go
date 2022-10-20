package services

import "service-api/app/core/category_menu/entities"

type CategoryMenuService interface {
	FindAll() ([]entities.CategoryMenu, error)
}
