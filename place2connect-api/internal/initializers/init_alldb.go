package initializers

import (
	"github.com/fxfrancky/place2connect-api/config"
	"gorm.io/gorm"
)

func LoadDatabases(conf *config.Config) *gorm.DB {

	db := ConnectDB(conf)
	ConnectRedis(conf)

	return db
}
