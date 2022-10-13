package entities

type UserType struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (UserType) TableName() string {
	return "user_type"
}
