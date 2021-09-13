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

	router.Static("/static/pictures", "./pictures")

	public := router.Group("/api/v1")
	{
		public.POST("/register", api.Register)
		public.POST("/login", api.Login)
		public.GET("/pictures", api.GetPictures)
		public.GET("/pictures/:id", api.GetPicture)
	}

	private := public.Group("/", middleware.LoginCheck())
	{
		private.GET("/logout", api.Logout)
		private.POST("/pictures", api.CreatePicture)
	}

	return router
}
