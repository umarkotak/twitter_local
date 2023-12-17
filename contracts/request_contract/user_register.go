package request_contract

type UserRegister struct {
	Name     string `json:"name" validate:"required" example:"jhone doe"`
	Username string `json:"username" validate:"required" example:"jhone_doe"`
	Password string `json:"password" validate:"required" example:"rahasia"`
	Email    string `json:"email" validate:"required" example:"jhone.doe@gmail.com"`
	PhotoUrl string `json:"photo_url" validate:"required" example:"https://picsum.photos/300/300"`
	About    string `json:"about" example:"hello world"`
}
