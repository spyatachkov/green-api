package validator

import (
	"regexp"

	"github.com/spyatachkov/green-api/backend/internal/models"
)

var phoneRegex = regexp.MustCompile(`^\d+$`)

func ValidatePhone(phone string) error {
	if phone == "" {
		return models.ErrPhoneEmpty
	}

	if !phoneRegex.MatchString(phone) {
		return models.ErrPhoneInvalidFormat
	}

	length := len(phone)
	if length < 10 || length > 15 {
		return models.ErrPhoneInvalidLength
	}

	return nil
}

func IsValidPhone(phone string) bool {
	return ValidatePhone(phone) == nil
}
