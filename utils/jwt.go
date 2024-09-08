package utils

import (
	"go-clean-architecture/module/user/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *model.User, secret string, expiry int) (string, error) {
	// Create a new token object, specifying signing method and claims
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Minute * time.Duration(expiry)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func CreateRefereshToken(user *model.User, secret string, expiry int) (string, error) {
	// Create a new token object, specifying signing method and claims
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString([]byte(secret))
}
