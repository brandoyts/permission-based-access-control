package unit

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	t.Parallel()
	token, err := jwtAuth.CreateToken(jwt.MapClaims{
		"sub": "tester",
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	t.Parallel()

	// Setup
	validToken, _ := jwtAuth.CreateToken(jwt.MapClaims{
		"sub": "tester",
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	})

	expiredToken, _ := jwtAuth.CreateToken(jwt.MapClaims{
		"sub": "tester",
		"exp": time.Now().Add(-10 * time.Minute).Unix(), // expired
	})

	tests := []struct {
		name        string
		token       string
		wantErr     bool
		expectedSub string
	}{
		{
			name:        "valid token",
			token:       validToken,
			wantErr:     false,
			expectedSub: "tester",
		},
		{
			name:    "expired token",
			token:   expiredToken,
			wantErr: true,
		},
		{
			name:    "invalid token format",
			token:   "invalid.token.here",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := jwtAuth.DecodeToken(tt.token)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
