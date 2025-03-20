package utils

import "fmt"

func FormatError(msg error, err error) error {
	return fmt.Errorf("%w: %w", msg, err)
}
