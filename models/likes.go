package models

import (
	"gorm.io/gorm"

	"github.com/yamachoo/media_back/db"
)

type Like struct {
	gorm.Model
	PictureId string `gorm:"not null"`
	UserId    uint   `gorm:"not null"`
}

func init() {
	db.AutoMigrateDB(&Like{})
}
