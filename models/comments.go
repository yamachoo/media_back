package models

import (
	"github.com/yamachoo/media_back/db"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PictureId string `gorm:"not null"`
	UserId    uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
}

func init() {
	db.AutoMigrateDB(&Comment{})
}
