package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ChechPassword(password, has string) bool {
	err := bycript.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
