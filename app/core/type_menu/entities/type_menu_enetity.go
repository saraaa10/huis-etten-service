package entities

import "time"

type TypeMenu struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (TypeMenu) TableName() string {
	return "type_menu"
}
