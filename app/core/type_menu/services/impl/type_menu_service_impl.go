package impl

import (
	"service-api/app/core/type_menu/entities"
	"service-api/app/core/type_menu/repositories"
)

type typeMenuService struct {
	repository repositories.TypeMenuRepository
}

func NewTypeMenuService(repo repositories.TypeMenuRepository) *typeMenuService {
	return &typeMenuService{
		repository: repo,
	}
}

func (service *typeMenuService) FindAll() ([]entities.TypeMenu, error) {
	typeMenus, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return typeMenus, nil
}
