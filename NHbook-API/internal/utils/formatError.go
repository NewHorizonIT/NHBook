package utils

import "fmt"

func FormatError(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	var err error

	for _, e := range errs {
		if err == nil {
			err = e
		} else {
			err = fmt.Errorf("%v: %v", err, e)
		}
	}

	return err
}
