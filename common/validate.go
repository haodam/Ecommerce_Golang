package common

import (
	"errors"
	"strings"
)

func ValidateString(str string) error {
	str = strings.TrimSpace(str)

	if str == "" {
		return errors.New("string can not be empty")
	}
	return nil
}
