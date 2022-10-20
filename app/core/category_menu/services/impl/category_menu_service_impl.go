package impl

import (
	"service-api/app/core/category_menu/entities"
	"service-api/app/core/category_menu/repositories"
)

type categoryMenuService struct {
	repo repositories.CategoryMenuRepository
}

func NewCategoryMenuService(repo repositories.CategoryMenuRepository) *categoryMenuService {
	return &categoryMenuService{
		repo: repo,
	}
}

func (service *categoryMenuService) FindAll() ([]entities.CategoryMenu, error) {
	categoryMenus, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return categoryMenus, nil
}
