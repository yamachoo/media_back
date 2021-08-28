package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/yamachoo/media_back/models"
)

type RegisterRequest struct {
	Name string `json:"name" binding:"required,max=24,min=1"`
	LoginRequest
}

type LoginRequest struct {
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

func Login(c *gin.Context) {
	var req LoginRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}

	user, err := models.GetUserByEmail(req.Email)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	session := sessions.Default(c)
	if session.Get("userId") == nil {
		session.Set("userId", user.ID)
		session.Save()
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Status(http.StatusOK)
}
