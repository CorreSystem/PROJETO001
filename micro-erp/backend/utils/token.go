package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(getSecret())

func getSecret() string {
	s := os.Geten("JWT_SECRET")
	if s == "" {
		s = "secret"
	}
	return s
}

func GenerateToken(userID uint, email, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":    userID,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokern.SignedString(secret)
}
