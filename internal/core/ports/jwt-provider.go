package ports

import "github.com/golang-jwt/jwt/v5"

type JwtProvider interface {
	CreateToken(claims jwt.MapClaims) (string, error)
	DecodeToken(token string) (jwt.MapClaims, error)
}
