package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(HashPassword), err
}

func CekPassword(hashPassword, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}