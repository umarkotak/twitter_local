package coba

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DaftarinPath(r *gin.Engine) {
	r.GET("/coba", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
