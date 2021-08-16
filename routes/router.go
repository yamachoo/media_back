package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yamachoo/media_back/api"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/resister", api.Register)
	}

	return router
}
