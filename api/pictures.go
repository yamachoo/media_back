package api

import (
	"encoding/base64"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yamachoo/media_back/models"
	"github.com/yamachoo/media_back/utils"
)

type CreatePictureRequest struct {
	Filename string `json:"filename" binding:"required,max=100,min=1"`
	Picture  string `json:"picture" binding:"required,min=1"`
}

func CreatePicture(c *gin.Context) {
	var req CreatePictureRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}

	var id string
	for exist := true; exist; {
		id, err = utils.MakeRandomStr(8)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		exist = models.CheckPictureById(id)
	}

	data, err := base64.StdEncoding.DecodeString(req.Picture)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	path := "pictures/" + id + ".jpg"
	err = os.WriteFile(path, data, os.ModePerm)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	session := sessions.Default(c)
	picture := models.Picture{
		ID:       id,
		UserId:   session.Get("userId").(uint),
		Filename: req.Filename,
		Path:     "/static/" + path,
	}
	err = picture.Create()
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func GetPictures(c *gin.Context) {
	pictures, err := models.GetPictures()
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, pictures)
}
