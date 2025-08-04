package unit

import (
	"testing"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
	"github.com/brandoyts/permission-based-access-control/internal/core/services"
	"github.com/brandoyts/permission-based-access-control/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupMockUserRepository(ctrl *gomock.Controller, email, password string) *mock.MockUserRepository {
	mockRepo := mock.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		FindByEmail(gomock.Any(), gomock.Any()).
		Return(domain.User{
			Email:    email,
			Password: password,
		}, nil)
	return mockRepo
}

func setupMockJwtProvider(ctrl *gomock.Controller, mockedToken string) *mock.MockJwtProvider {
	mockProvider := mock.NewMockJwtProvider(ctrl)
	mockProvider.EXPECT().
		CreateToken(gomock.Any()).
		Return(mockedToken, nil)
	return mockProvider
}

func TestLogin(t *testing.T) {
	const (
		testEmail    = "test@mail.com"
		testPassword = "secret"
		testToken    = "token123"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := setupMockUserRepository(ctrl, testEmail, testPassword)
	jwtProvider := setupMockJwtProvider(ctrl, testToken)

	authService := services.NewAuthService(userRepo, jwtProvider)

	token, err := authService.Login(testEmail, testPassword)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assert.NotEmpty(t, token)
}
