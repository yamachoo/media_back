package models

import (
	"gorm.io/gorm"

	"github.com/yamachoo/media_back/db"
)

type User struct {
	gorm.Model
	Name     string    `gorm:"not null" binding:"required,max=24,min=1"`
	Email    string    `gorm:"unique;not null" binding:"required,email,max=100"`
	Password string    `gorm:"not null"`
	Pictures []Picture `gorm:"foreignKey:UserId" binding:"-"`
	Likes    []Like    `gorm:"foreignKey:UserId" binding:"-"`
	Comments []Comment `gorm:"foreignKey:UserId" binding:"-"`
}

func init() {
	db.AutoMigrateDB(&User{})
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	return db.Validate.Struct(u)
}

func (u *User) Create() error {
	db := db.GetDB()
	return db.Create(u).Error
}
