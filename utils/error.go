package utils

import (
	"errors"

	"gorm.io/gorm"
)

type CustomError struct {
	Code int
	Msg  error
}

func ErrorWrap(err error) CustomError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return CustomError{Code: 404, Msg: err}
	}

	return CustomError{Code: 500, Msg: err}
}
