package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/yamachoo/media_back/models"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,max=24,min=1"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,max=24,min=8"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}
	err = user.Create()
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
