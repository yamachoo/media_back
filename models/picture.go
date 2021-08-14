package models

import (
	"time"

	"gorm.io/gorm"
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
