package models

import (
	"database/sql"
	"time"
)

type (
	Post struct {
		ID        int64        `db:"id"`
		Content   string       `db:"content"`
		UserID    string       `db:"user_id"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt time.Time    `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
	}
)
