package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")
		if userId == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}
		c.Next()
	}
}
