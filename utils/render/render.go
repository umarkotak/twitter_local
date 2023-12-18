package render

import (
	"github.com/gin-gonic/gin"
	"github.com/umarkotak/twitter_local/error_const"
)

func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	body := map[string]interface{}{
		"data":    data,
		"success": true,
	}
	c.JSON(200, body)
}

func Error(c *gin.Context, err error, customMessage string) {
	errObj, found := error_const.ERR_MAP[err]
	if !found {
		errObj = error_const.ERR_MAP[error_const.INTERNAL_SERVER_ERROR]
	}

	message := errObj.Message
	if customMessage != "" {
		message = customMessage
	}

	body := map[string]interface{}{
		"error": map[string]interface{}{
			"code":     errObj.Code,
			"title":    errObj.Title,
			"message":  message,
			"internal": err.Error(),
		},
		"success": true,
	}
	c.JSON(errObj.StatusCode, body)
}
