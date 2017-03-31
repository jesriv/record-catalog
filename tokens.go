package main

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "machines making sounds"
)

func NewToken(user *User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["user_id"]	= user.ID
	claims["username"]	= user.Username
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(signingKey))

	return tokenString, err
}

func ValidToken(myToken string, myKey string) bool {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})

	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}