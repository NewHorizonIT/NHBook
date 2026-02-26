package errs

import "errors"

var (
	ErrorNotFound       = errors.New("not found")
	ErrorUnauthorized   = errors.New("unauthorized")
	ErrorInvalidInput   = errors.New("invalid input")
	ErrorInternalServer = errors.New("internal server error")
)
