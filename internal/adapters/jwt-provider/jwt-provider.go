package jwtProvider

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JwtProvider struct {
	secretKey []byte
}

const ErrInvalidToken = "invalid token"

func New(secretKey string) *JwtProvider {
	return &JwtProvider{
		secretKey: []byte(secretKey),
	}
}

func (jp *JwtProvider) CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jp.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jp *JwtProvider) DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jp.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New(ErrInvalidToken)
	}

	return token.Claims.(jwt.MapClaims), nil
}
