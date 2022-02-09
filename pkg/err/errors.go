package errs

import "errors"

var (
	ErrDb              = errors.New("unexpected database error")
	ErrNoRowsFound     = errors.New("no values found")
	ErrServerErr       = errors.New("internal server error")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidToken    = errors.New("invalid token")
	ErrDuplicateValue  = errors.New("this value already exists")
	ErrEmailMissing    = errors.New("email is missing")
	ErrInvalidEmail    = errors.New("invalid email")
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}
