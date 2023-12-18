package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/contracts/request_contract"
	"github.com/umarkotak/twitter_local/services/user_service"
	"github.com/umarkotak/twitter_local/utils/jwt_auth"
	"github.com/umarkotak/twitter_local/utils/render"
)

// ShowAccount godoc
//
//	@Summary	Register
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		user	body		request_contract.UserRegister	true	"body"
//	@Success	200		{object}	interface{}
//	@Router		/users/register [post]
func Register(c *gin.Context) {
	params := request_contract.UserRegister{}
	err := c.BindJSON(&params)
	if err != nil {
		logrus.WithContext(c).Error(err)
		c.JSON(400, map[string]interface{}{"error": err.Error()})
		return
	}

	err = user_service.Register(c, params)
	if err != nil {
		logrus.WithContext(c).Error(err)

		if _, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(500, map[string]interface{}{"error": err.Error()})
		return
	}

	render.Success(c, nil)
}

func Login(c *gin.Context) {
	params := request_contract.UserLogin{}
	err := c.BindJSON(&params)
	if err != nil {
		logrus.WithContext(c).Error(err)
		render.Error(c, err, "")
		return
	}

	response, err := user_service.Login(c, params)
	if err != nil {
		logrus.WithContext(c).Error(err)
		render.Error(c, err, "")
		return
	}

	render.Success(c, response)
}

func MyProfile(c *gin.Context) {
	user, err := jwt_auth.DecodeUserAuthToken(c, c.GetHeader("Authorization"))
	if err != nil {
		logrus.WithContext(c).Error(err)
		render.Error(c, err, "")
		return
	}

	response, err := user_service.GetByID(c, user.ID)
	if err != nil {
		logrus.WithContext(c).Error(err)
		render.Error(c, err, "")
		return
	}

	render.Success(c, response)
}

func ProfileByUsername(c *gin.Context) {
	response, err := user_service.GetByUsername(c, c.Param("username"))
	if err != nil {
		logrus.WithContext(c).Error(err)
		render.Error(c, err, "")
		return
	}

	render.Success(c, response)
}
