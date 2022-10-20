package entities

import "time"

type UserType struct {
	ID        int       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`
}

func (UserType) TableName() string {
	return "user_type"
}
