package impl

import (
	"service-api/app/core/menu/entities"

	"gorm.io/gorm"
)

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) FindAll() ([]entities.Menu, error) {
	var menus []entities.Menu
	err := r.db.Preload("Type").Find(&menus).Error
	if err != nil {
		return nil, err
	}

	return menus, nil
}
