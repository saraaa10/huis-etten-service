package impl

import (
	"service-api/app/core/menu/entities"
	"service-api/app/core/menu/repositories"
)

type menuService struct {
	menuRepo repositories.MenuRepository
}

func NewMenuService(repo repositories.MenuRepository) *menuService {
	return &menuService{menuRepo: repo}
} 

func (s *menuService) FindAll() ([]entities.Menu, error) {
	menu, err := s.menuRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return menu, nil
}