package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yamachoo/media_back/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthForm struct {
	Name     string `json:"name" binding:"required,max=24,min=1"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,max=24,min=8"`
}

func Register(c *gin.Context) {
	var form AuthForm
	err := c.BindJSON(&form)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	user.Password, err = bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = user.Create()
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
