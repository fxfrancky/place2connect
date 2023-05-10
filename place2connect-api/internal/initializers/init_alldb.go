package initializers

import (
	"log"

	"github.com/fxfrancky/place2connect-api/config"
	"gorm.io/gorm"
)

func LoadDatabases(path string) *gorm.DB {
	conf, err := config.LoadConfig(path)
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	ConnectRedis(&conf)
	db := ConnectDB(&conf)
	return db
}
