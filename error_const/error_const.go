package error_const

import "fmt"

var (
	INTERNAL_SERVER_ERROR = fmt.Errorf("internal_server_error")
	BAD_REQUEST           = fmt.Errorf("bad_request")
	INVALID_AUTH          = fmt.Errorf("invalid_auth")
	UNAUTHORIZED          = fmt.Errorf("unauthorized")
)

type ErrorObj struct {
	StatusCode int
	Code       string
	Title      string
	Message    string
}

var (
	ERR_MAP = map[error]ErrorObj{
		INTERNAL_SERVER_ERROR: {
			StatusCode: 500,
			Code:       "internal_server_error",
			Title:      "internal server error",
			Message:    "internal server error",
		},
		BAD_REQUEST: {
			StatusCode: 400,
			Code:       "bad_request",
			Title:      "bad request",
			Message:    "bad request",
		},
		INVALID_AUTH: {
			StatusCode: 400,
			Code:       "invalid_auth",
			Title:      "invalid auth",
			Message:    "bad request",
		},
		UNAUTHORIZED: {
			StatusCode: 401,
			Code:       "unauthorized",
			Title:      "unauthorized",
			Message:    "unauthorized",
		},
	}
)
