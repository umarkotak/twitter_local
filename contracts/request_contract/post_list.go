package request_contract

type (
	PostList struct {
		UserID int64 `json:"user_id" validate:"required"`
	}
)
