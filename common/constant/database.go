package constant

import (
	"fmt"
	"log"

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
