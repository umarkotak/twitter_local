package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
