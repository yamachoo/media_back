package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/yamachoo/media_back/config"
)

var db *gorm.DB

func init() {
	c := config.GetConfig()
	var err error
	db, err = gorm.Open(sqlite.Open(c.GetString("db.url")))
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrateDB(models ...interface{}) {
	db.AutoMigrate(models...)
}
