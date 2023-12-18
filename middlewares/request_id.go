package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/umarkotak/twitter_local/models"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		commonReqInterface, found := c.Get("common_request")
		commonReq := models.CommonRequest{}
		if found {
			commonReq = commonReqInterface.(models.CommonRequest)
		}

		reqID, _ := uuid.NewRandom()
		commonReq.RequestID = reqID.String()
		c.Set("common_request", commonReq)

		c.Next()
	}
}
