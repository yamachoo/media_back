package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/yamachoo/media_back/db"
)

type Picture struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserId    uint           `gorm:"not null"`
	Filename  string         `gorm:"not null"`
	Picture   []byte         `gorm:"not null"`
	Likes     []Like         `gorm:"foreignKey:PictureId"`
	Comments  []Comment      `gorm:"foreignKey:PictureId"`
}

func init() {
	db.AutoMigrateDB(&Picture{})
}
