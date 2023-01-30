package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func ValidatePassword(plain string, userPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(plain))

	return err == nil
}
