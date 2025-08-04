package services

import (
	"context"
	"errors"
	"time"

	"github.com/brandoyts/permission-based-access-control/internal/core/ports"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userRepository ports.UserRepository
	jwtProvider    ports.JwtProvider
}

const ErrInvalidCredentials = "invalid credentials"

func NewAuthService(userRepository ports.UserRepository, jwtProvider ports.JwtProvider) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtProvider:    jwtProvider,
	}
}

func (as *AuthService) Login(email string, password string) (string, error) {
	user, err := as.userRepository.FindByEmail(context.Background(), email)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New(ErrInvalidCredentials)
	}

	token, err := as.jwtProvider.CreateToken(jwt.MapClaims{
		"sub":         user.Email,
		"exp":         time.Now().Add(10 * time.Minute).Unix(),
		"permissions": user.Permissions,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
