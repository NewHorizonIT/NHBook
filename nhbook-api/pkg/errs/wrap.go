package errs

import "fmt"

type AppError struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Original string `json:"-"`
	Err      error  `json:"-"`
	Cause    error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s", e.Code)
	}
	return e.Message
}

func Wrap(err error, original string, code, message string) *AppError {
	return &AppError{
		Code:     code,
		Message:  message,
		Original: original,
		Err:      err,
	}
}

func (e *AppError) WithCause(cause error) *AppError {
	return &AppError{
		Code:     e.Code,
		Message:  e.Message,
		Original: e.Original,
		Err:      e.Err,
		Cause:    cause,
	}
}
