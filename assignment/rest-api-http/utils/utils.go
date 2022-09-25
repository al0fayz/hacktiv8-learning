package utils

import "errors"

func ShowError(message string) error {
	return errors.New(message)
}
