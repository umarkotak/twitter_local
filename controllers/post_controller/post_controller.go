package post_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/contracts/request_contract"
	"github.com/umarkotak/twitter_local/error_const"
	"github.com/umarkotak/twitter_local/models"
	"github.com/umarkotak/twitter_local/services/post_service"
	"github.com/umarkotak/twitter_local/utils/render"
)

func Create(c *gin.Context) {
	params := request_contract.PostCreate{}
	err := c.BindJSON(&params)
	if err != nil {
		logrus.WithContext(c).Error(err)
		render.Error(c, err, "")
		return
	}

	commonReqInterface, _ := c.Get("common_request")
	commonReq, ok := commonReqInterface.(models.CommonRequest)
	if !ok {
		render.Error(c, error_const.INVALID_COMMON_REQUEST, "")
		return
	}

	params.UserID = commonReq.User.ID

	response, err := post_service.Create(c, params)
	if err != nil {
		logrus.WithContext(c).Error(err)
		render.Error(c, err, "")
		return
	}

	render.Success(c, response)
}

func GetMyPosts(c *gin.Context) {
	// Common flow pembuatan API
	// 1. Buat repository layer nya - select from posts where ...
	// 2. Buat service layer nya
	// 3. Buat controller layer nya
	commonReqInterface, _ := c.Get("common_request")
	commonReq, ok := commonReqInterface.(models.CommonRequest)
	if !ok {
		render.Error(c, error_const.INVALID_COMMON_REQUEST, "")
		return
	}

	params := request_contract.PostList{
		UserID: commonReq.User.ID,
	}

	response, err := post_service.GetByUserID(c, params)
	if err != nil {
		logrus.WithContext(c).Error(err)
		render.Error(c, err, "")
		return
	}

	render.Success(c, response)
}

func GetList(c *gin.Context) {

}

func GetDetail(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
