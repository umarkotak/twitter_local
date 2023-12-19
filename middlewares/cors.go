package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/umarkotak/twitter_local/utils/render"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Animapu-User-Uid, X-Visitor-Id, X-From-Path",
		)

		if c.Request.Method == "OPTIONS" {
			render.Success(c, nil)
			return
		}

		c.Next()
	}
}
