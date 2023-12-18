package post_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/contracts/request_contract"
	"github.com/umarkotak/twitter_local/contracts/response_contract"
	"github.com/umarkotak/twitter_local/models"
	"github.com/umarkotak/twitter_local/repositories/post_repository"
)

func Create(ctx context.Context, params request_contract.PostCreate) (response_contract.PostCreate, error) {
	err := validator.New().Struct(params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return response_contract.PostCreate{}, err
	}

	id, err := post_repository.Create(ctx, models.Post{
		Content: params.Content,
		UserID:  params.UserID,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return response_contract.PostCreate{}, err
	}

	return response_contract.PostCreate{
		ID: id,
	}, nil
}
