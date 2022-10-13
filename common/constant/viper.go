package constant

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseStruct struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASS"`
	Database string `mapstructure:"DB_NAME"`
}

func LoadEnv(path string) (db DatabaseStruct, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err.Error())
	}

	err = viper.Unmarshal(&db)
	if err != nil {
		log.Fatal("Error unmarshal .env file")
		log.Fatal(err.Error())
	}

	return
}
