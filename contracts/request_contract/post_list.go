package request_contract

type (
	PostList struct {
		UserID string `json:"user_id" validate:"required"`
	}
)
