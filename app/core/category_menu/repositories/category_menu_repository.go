package repositories

import "service-api/app/core/category_menu/entities"

type CategoryMenuRepository interface {
	FindAll() ([]entities.CategoryMenu, error)
}
