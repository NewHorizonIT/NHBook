package validator

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetErrorValidate(err error) map[string]string {
	var errs validator.ValidationErrors

	errorsMap := make(map[string]string)
	if !errors.As(err, &errs) {
		return errorsMap
	}

	for _, e := range errs {
		field := strings.ToLower(e.Field())
		tag := e.Tag()

		switch tag {
		case "required":
			errorsMap[field] = "field is required"
		case "email":
			errorsMap[field] = "invalid email format"
		case "min":
			errorsMap[field] = "value too short"
		case "max":
			errorsMap[field] = "value too long"
		default:
			errorsMap[field] = "invalid value"
		}
	}

	return errorsMap
}
