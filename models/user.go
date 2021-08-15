package models

import (
	"gorm.io/gorm"

	"github.com/yamachoo/media_back/db"
)

type User struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Pictures []Picture `gorm:"foreignKey:UserId"`
	Likes    []Like    `gorm:"foreignKey:UserId"`
	Comments []Comment `gorm:"foreignKey:UserId"`
}

func init() {
	db.AutoMigrateDB(&User{})
}
