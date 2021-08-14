package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Pictures []Picture `gorm:"foreignKey:UserId"`
	Likes    []Like    `gorm:"foreignKey:UserId"`
	Comments []Comment `gorm:"foreignKey:UserId"`
}
