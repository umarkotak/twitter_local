package post_repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/config"
	"github.com/umarkotak/twitter_local/models"
)

func Create(ctx context.Context, post models.Post) (int64, error) {
	rows, err := config.Get().MasterDB.NamedQueryContext(ctx, `
		INSERT INTO posts (
			content, user_id, created_at, updated_at
		) VALUES (
			:content, :user_id, NOW(), NOW()
		) RETURNING id
	`, post)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return 0, err
	}

	id := int64(0)

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return 0, err
		}
	}

	return id, nil
}
