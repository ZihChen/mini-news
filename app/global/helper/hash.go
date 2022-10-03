package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (value string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return string(bytes), err
	}

	return string(bytes), err
}
