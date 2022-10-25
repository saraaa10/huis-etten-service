package entities

import (
	categoryMenu "service-api/app/core/category_menu/entities"
	typeMenu "service-api/app/core/type_menu/entities"
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID          uint                      `gorm:"primaryKey;column:id;autoIncrement;not null" json:"id"`
	Name        string                    `gorm:"column:name;not null" json:"name"`
	Description string                    `gorm:"column:description;not null;" json:"description"`
	CategoryId  int                       `json:"-"`
	Category    categoryMenu.CategoryMenu `gorm:"foreignKey:CategoryId" json:"category"`
	TypeId      int                       `json:"-"`
	Type        typeMenu.TypeMenu         `gorm:"foreignKey:TypeId" json:"type"`
	Price       int                       `gorm:"column:price;not null;" json:"price"`
	CreatedAt   time.Time                 `gorm:"default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
	UpdatedAt   time.Time                 `gorm:"default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt            `gorm:"column:deleted_at" json:"-"`
}

func (Menu) TableName() string {
	return "menu"
}
