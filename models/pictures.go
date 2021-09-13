package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/yamachoo/media_back/db"
)

type Picture struct {
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	UserId    uint           `json:"userId" gorm:"not null" binding:"required"`
	Filename  string         `json:"filename" gorm:"not null" binding:"required,max=100,min=1"`
	Path      string         `json:"path" gorm:"not null" binding:"required"`
	Likes     []Like         `json:"likes" gorm:"foreignKey:PictureId" binding:"-"`
	Comments  []Comment      `json:"comments" gorm:"foreignKey:PictureId" binding:"-"`
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

func CheckPictureById(id string) bool {
	var picture Picture
	db := db.GetDB()
	if result := db.Where(&Picture{ID: id}).First(&picture); result.Error != nil {
		return false
	}
	return true
}

func GetPictures() ([]Picture, error) {
	var pictures []Picture
	db := db.GetDB()
	result := db.Find(&pictures)
	return pictures, result.Error
}

func GetPicture(id string) (Picture, error) {
	var picture Picture
	db := db.GetDB()
	result := db.Where(&Picture{ID: id}).First(&picture)
	return picture, result.Error
}
