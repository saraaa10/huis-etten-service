package entities

import (
	"service-api/app/core/user_type/entities"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint              `gorm:"primaryKey; autoIncrement; not null" json:"id"`
	Name       string            `gorm:"column:name;not null" json:"name"`
	UserTypeId uint              `json:"-"`
	UserType   entities.UserType `gorm:"foreignKey:UserTypeId" json:"user_type"`
	Username   string            `gorm:"column:username" json:"username"`
	Email      string            `gorm:"column:email; not null" json:"email"`
	Password   string            `gorm:"column:password; not null" json:"password"`
	CreatedAt  time.Time         `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time         `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt    `gorm:"column:deleted_at" json:"-"`
}

func (User) TableName() string {
	return "users"
}
