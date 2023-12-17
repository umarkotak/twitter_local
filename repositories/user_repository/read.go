package user_repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/config"
	"github.com/umarkotak/twitter_local/models"
)

func GetByUsername(ctx context.Context, username string) (models.User, error) {
	user := models.User{}

	err := config.Get().MasterDB.GetContext(ctx, &user, `
		SELECT
			id,
			name,
			username,
			password,
			email,
			photo_url,
			about,
			created_at,
			updated_at,
			deleted_at
		FROM users
		WHERE
			username = $1
			AND deleted_at IS NULL
	`, username)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return models.User{}, err
	}

	return user, nil
}

func GetByID(ctx context.Context, id string) (models.User, error) {
	user := models.User{}

	err := config.Get().MasterDB.GetContext(ctx, &user, `
		SELECT
			id,
			name,
			username,
			password,
			email,
			photo_url,
			about,
			created_at,
			updated_at,
			deleted_at
		FROM users
		WHERE
			id = $1
			AND deleted_at IS NULL
	`, id)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return models.User{}, err
	}

	return user, nil
}
