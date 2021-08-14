package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	PictureId string `gorm:"not null"`
	UserId    uint   `gorm:"not null"`
}
