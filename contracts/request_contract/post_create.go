package request_contract

type (
	PostCreate struct {
		Content string `json:"content" validate:"required"`
		UserID  int64  `json:"user_id" validate:"required"`
	}
)
