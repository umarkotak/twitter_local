package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/umarkotak/twitter_local/error_const"
	"github.com/umarkotak/twitter_local/models"
)

func GetCommonRequest(c *gin.Context) (models.CommonRequest, error) {
	commonReqInterface, _ := c.Get("common_request")
	commonReq, ok := commonReqInterface.(models.CommonRequest)
	if !ok {
		return models.CommonRequest{}, error_const.INVALID_COMMON_REQUEST
	}
	return commonReq, nil
}
