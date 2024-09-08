package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
)

func HashPassword(password string) (string, error) {

	costOfPassword := 8

	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, costOfPassword)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	return err == nil
}
