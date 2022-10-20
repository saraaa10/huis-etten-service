package entities

type CategoryMenu struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (CategoryMenu) TableName() string {
	return "category_menu"
}
