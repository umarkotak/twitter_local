package response_contract

import "time"

type (
	PostList struct {
		Posts []Post `json:"posts"`
	}

	Post struct {
		ID             int64     `json:"id"`
		Content        string    `json:"content"`
		CreatedAt      time.Time `json:"created_at"`
		AuthorID       string    `json:"author_id"`
		AuthorName     string    `json:"author_name"`
		AuthorUsername string    `json:"author_username"`
	}
)
