package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/yamachoo/media_back/api"
	"github.com/yamachoo/media_back/config"
)

func SetupRouter() *gin.Engine {
	c := config.GetConfig()
	router := gin.Default()
	store := cookie.NewStore([]byte(c.GetString("router.cookie")))
	router.Use(sessions.Sessions(c.GetString("router.session"), store))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/resister", api.Register)
		v1.POST("/login", api.Login)
	}

	return router
}
