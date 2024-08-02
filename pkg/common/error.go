package common

import "errors"

var (
	ErrPasswordDoesntMatched = errors.New("password doesn't matched")
	ErrEmailAlreadyUsed      = errors.New("email already used")
	ErrRecordNotFound        = errors.New("record not found")
)
