package response_contract

import (
	"time"
)

type (
	UserMyProfile struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Email     string    `json:"email"`
		PhotoUrl  string    `json:"photo_url"`
		About     string    `json:"about"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
