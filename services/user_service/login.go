package user_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/contracts/request_contract"
	"github.com/umarkotak/twitter_local/contracts/response_contract"
	"github.com/umarkotak/twitter_local/error_const"
	"github.com/umarkotak/twitter_local/repositories/user_repository"
	"github.com/umarkotak/twitter_local/utils/jwt_auth"
)

func Login(ctx context.Context, params request_contract.UserLogin) (response_contract.UserAuthToken, error) {
	err := validator.New().Struct(params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return response_contract.UserAuthToken{}, err
	}

	user, err := user_repository.GetByUsername(ctx, params.Username)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return response_contract.UserAuthToken{}, err
	}

	if user.Password != params.Password {
		err = error_const.BAD_REQUEST
		return response_contract.UserAuthToken{}, err
	}

	token, err := jwt_auth.GenUserAuthToken(ctx, user)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return response_contract.UserAuthToken{}, err
	}

	return response_contract.UserAuthToken{
		Token: token,
	}, nil
}
