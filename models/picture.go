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
	UserId    uint           `gorm:"not null" binding:"required"`
	Filename  string         `gorm:"not null" binding:"required,max=100,min=1"`
	Path      string         `gorm:"not null" binding:"required"`
	Likes     []Like         `gorm:"foreignKey:PictureId" binding:"-"`
	Comments  []Comment      `gorm:"foreignKey:PictureId" binding:"-"`
}

func init() {
	db.AutoMigrateDB(&Picture{})
}

func (p *Picture) BeforeSave(tx *gorm.DB) error {
	return db.Validate.Struct(p)
}

func (p *Picture) Create() error {
	db := db.GetDB()
	return db.Create(p).Error
}
