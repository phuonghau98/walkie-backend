package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaim struct {
	UserInfo struct {
		ID string
		Email string
	}
	jwt.StandardClaims
}

func EncodeToken(user *User) (string, error) {
	expirationTime := time.Now().Add(300 * time.Hour).Unix()
	claims := TokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    "go.micro.srv.user",
		},
	}

	claims.UserInfo.Email = user.Email
	claims.UserInfo.ID = user.ID.String()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func DecodeToken (token string) (*TokenClaim, error) {
	decodedClaims := &TokenClaim{}
	_, err := jwt.ParseWithClaims(token, decodedClaims, func (token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return decodedClaims, nil
}