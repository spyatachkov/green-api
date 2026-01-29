package models

import "errors"

var (
	ErrPhoneEmpty         = errors.New("phone number cannot be empty")
	ErrPhoneInvalidFormat = errors.New("phone number must contain only digits")
	ErrPhoneInvalidLength = errors.New("phone number must be between 10 and 15 digits")
)
