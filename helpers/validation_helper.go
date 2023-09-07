package helpers

import (
	"errors"
	"regexp"
)

func ValidateTextLength(value string, minLength int) (bool, error) {
	if len(value) < minLength {
		return false, errors.New("text length is too short")
	}
	return true, nil
}

func ValidateTextLengthRange(value string, minLength int, maxLength int) (bool, error) {
	if len(value) < minLength || len(value) > maxLength {
		return false, errors.New("text length is too short or too long")
	}
	return true, nil
}

func ValidateEmail(value string) (bool, error) {
	regexp := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	if !regexp.MatchString(value) {
		return false, errors.New("invalid email format")
	}
	return true, nil
}

func ValidatePassword(value string) (bool, error) {
	if len(value) < 8 {
		return false, errors.New("password is too short")
	}
	return true, nil
}

func ValidateExistance(value string) (bool, error) {
	if len(value) == 0 {
		return false, errors.New("value does not exist")
	}
	return true, nil
}
