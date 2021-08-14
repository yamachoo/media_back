package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/yamachoo/media_back/config"
	"github.com/yamachoo/media_back/models"
)

var db *gorm.DB

func init() {
	c := config.GetConfig()
	var err error
	db, err = gorm.Open(sqlite.Open(c.GetString("db.url")))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Picture{}, &models.Like{}, &models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
