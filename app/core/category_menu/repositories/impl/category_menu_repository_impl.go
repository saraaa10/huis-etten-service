package impl

import (
	"service-api/app/core/category_menu/entities"

	"gorm.io/gorm"
)

type categoryMenuRepository struct {
	db *gorm.DB
}

func NewCategoryMenuRepository(database *gorm.DB) *categoryMenuRepository {
	return &categoryMenuRepository{
		db: database,
	}
}

func (repo *categoryMenuRepository) FindAll() ([]entities.CategoryMenu, error) {
	var categoryMenus []entities.CategoryMenu

	err := repo.db.Find(&categoryMenus).Error
	if err != nil {
		return nil, err
	}

	return categoryMenus, nil
}
