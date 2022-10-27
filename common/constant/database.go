package constant

import (
	"fmt"
	"log"
	entitiesUser "service-api/app/core/users/entities"
	entitiesUserType "service-api/app/core/user_type/entities"
	entitiesTypeMenu "service-api/app/core/type_menu/entities"
	entitiesCategoryMenu "service-api/app/core/category_menu/entities"
	entitiesMenu "service-api/app/core/menu/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	env, err := LoadEnv("../..")
	if err != nil {
		log.Fatal(err.Error())
	}

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", env.Host, env.Port, env.Username, env.Password, env.Database)
	dsn, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed Connect to Database")
		log.Fatal(err.Error())
	}

	log.Println("Success Connect to Database")

	return dsn
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entitiesUser.User{},
		&entitiesUserType.UserType{},
		&entitiesTypeMenu.TypeMenu{},
		&entitiesCategoryMenu.CategoryMenu{},
		&entitiesMenu.Menu{},
	)
}
