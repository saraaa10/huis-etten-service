package impl

import (
	"service-api/app/core/type_menu/entities"

	"gorm.io/gorm"
)

type typeMenuRepository struct {
	db *gorm.DB
}

func NewTypeMenuRepository(db *gorm.DB) *typeMenuRepository {
	return &typeMenuRepository{db: db}
}

func (r *typeMenuRepository) FindAll() ([]entities.TypeMenu, error) {
	var typeMenus []entities.TypeMenu

	err := r.db.Find(&typeMenus).Error
	if err != nil {
		return nil, err
	}

	return typeMenus, nil
}
