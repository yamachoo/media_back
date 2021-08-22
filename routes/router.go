package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/yamachoo/media_back/api"
	"github.com/yamachoo/media_back/config"
	"github.com/yamachoo/media_back/middleware"
)

func SetupRouter() *gin.Engine {
	c := config.GetConfig()
	router := gin.Default()
	store := cookie.NewStore([]byte(c.GetString("router.cookie")))
	router.Use(sessions.Sessions(c.GetString("router.session"), store))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	open := router.Group("/api/v1")
	{
		open.POST("/register", api.Register)
		open.POST("/login", api.Login)
	}

	auth := open.Group("/", middleware.LoginCheck())
	{
		auth.GET("/logout", api.Logout)
	}

	return router
}
