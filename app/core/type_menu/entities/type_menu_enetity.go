package entities

import (
	"time"

	"gorm.io/gorm"
)

type TypeMenu struct {
	ID        uint           `gorm:"primaryKey;column:id;autoIncrement;not null" json:"id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

func (TypeMenu) TableName() string {
	return "type_menu"
}
