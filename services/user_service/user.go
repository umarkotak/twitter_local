package user_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/contracts/response_contract"
	"github.com/umarkotak/twitter_local/repositories/user_repository"
)

func GetByID(ctx context.Context, userID int64) (response_contract.UserMyProfile, error) {
	user, err := user_repository.GetByID(ctx, userID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return response_contract.UserMyProfile{}, err
	}

	return response_contract.UserMyProfile{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		PhotoUrl:  user.PhotoUrl,
		About:     user.About.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func GetByUsername(ctx context.Context, username string) (response_contract.UserMyProfile, error) {
	user, err := user_repository.GetByUsername(ctx, username)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return response_contract.UserMyProfile{}, err
	}

	return response_contract.UserMyProfile{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		PhotoUrl:  user.PhotoUrl,
		About:     user.About.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
