package user_repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/config"
	"github.com/umarkotak/twitter_local/models"
)

func Create(ctx context.Context, user models.User) (string, error) {
	rows, err := config.Get().MasterDB.NamedQueryContext(ctx, `
		INSERT INTO users (
			name, username, password, email, photo_url, about, created_at, updated_at
		) VALUES (
			:name, :username, :password, :email, :photo_url, :about, NOW(), NOW()
		)
	`, user)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	id := ""

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return "", err
		}
	}

	return id, nil
}
