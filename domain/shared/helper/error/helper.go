package error

import "errors"

func SetErrorMessage(message string) (err error) {

	err = errors.New(message + " can not be empty...")

	return
}
