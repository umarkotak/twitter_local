package request_contract

type (
	PostCreate struct {
		Content string `json:"content" validate:"required"`
		UserID  string `json:"user_id" validate:"required"`
	}
)
