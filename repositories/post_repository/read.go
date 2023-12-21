package post_repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/config"
	"github.com/umarkotak/twitter_local/contracts/response_contract"
)

// Common golang DB commands

// config.Get().MasterDB.QueryContext() => rows
// config.Get().MasterDB.QueryRowContext() => row

// config.Get().MasterDB.SelectContext() => destination is a slice
// config.Get().MasterDB.GetContext() => one destination

// config.Get().MasterDB.ExecContext() => result

// ===

// config.Get().MasterDB.NamedQueryContext() => rows

// config.Get().MasterDB.NamedExecContext() => result

func GetByUserID(ctx context.Context, userID int64) ([]response_contract.Post, error) {
	rows, err := config.Get().MasterDB.QueryContext(ctx, `
		SELECT
			p.id,
			p.content,
			p.created_at,
			u.id AS author_id,
			u.name AS author_name,
			u.username AS author_username
		FROM posts p
		INNER JOIN users u ON u.id = p.user_id
		WHERE
			u.id = $1
		ORDER BY p.id DESC
	`, userID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return []response_contract.Post{}, err
	}

	posts := []response_contract.Post{}

	defer rows.Close()
	for rows.Next() {
		post := response_contract.Post{}
		err = rows.Scan(
			&post.ID,
			&post.Content,
			&post.CreatedAt,
			&post.AuthorID,
			&post.AuthorName,
			&post.AuthorUsername,
		)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return []response_contract.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
