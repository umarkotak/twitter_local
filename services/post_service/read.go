package post_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/contracts/request_contract"
	"github.com/umarkotak/twitter_local/contracts/response_contract"
	"github.com/umarkotak/twitter_local/repositories/post_repository"
)

func GetByUserID(ctx context.Context, params request_contract.PostList) (response_contract.PostList, error) {
	err := validator.New().Struct(params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return response_contract.PostList{}, err
	}

	posts, err := post_repository.GetByUserID(ctx, params.UserID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return response_contract.PostList{}, err
	}

	return response_contract.PostList{
		Posts: posts,
	}, nil
}
