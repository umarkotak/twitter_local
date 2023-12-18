package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/error_const"
	"github.com/umarkotak/twitter_local/models"
	"github.com/umarkotak/twitter_local/utils/jwt_auth"
	"github.com/umarkotak/twitter_local/utils/render"
)

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		commonReqInterface, found := c.Get("common_request")
		commonReq := models.CommonRequest{}
		if found {
			commonReq = commonReqInterface.(models.CommonRequest)
		}

		user, err := jwt_auth.DecodeUserAuthToken(c, c.GetHeader("Authorization"))
		if err != nil {
			logrus.WithContext(c).Error(err)
			render.Error(c, error_const.UNAUTHORIZED, err.Error())
			c.Abort()
			return
		}

		commonReq.User = user
		c.Set("common_request", commonReq)

		c.Next()
	}
}
