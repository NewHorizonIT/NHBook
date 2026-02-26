package errs

import "errors"

func MapError(err error) (int, string) {
	switch {
	case errors.Is(err, ErrorNotFound):
		return 404, "Resource not found"
	case errors.Is(err, ErrorUnauthorized):
		return 401, "Unauthorized access"
	case errors.Is(err, ErrorInvalidInput):
		return 400, "Invalid input provided"
	default:
		return 500, "Internal server error"
	}
}
