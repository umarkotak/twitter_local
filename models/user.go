package models

import (
	"database/sql"
	"time"
)

type (
	User struct {
		ID        int64          `db:"id"`
		Name      string         `db:"name"`
		Username  string         `db:"username"`
		Password  string         `db:"password"`
		Email     string         `db:"email"`
		PhotoUrl  string         `db:"photo_url"`
		About     sql.NullString `db:"about"`
		CreatedAt time.Time      `db:"created_at"`
		UpdatedAt time.Time      `db:"updated_at"`
		DeletedAt sql.NullTime   `db:"deleted_at"`
	}
)

func tes() {
}
