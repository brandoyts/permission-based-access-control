package unit

import (
	"context"
	"testing"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
)

func TestUserRepositoryCreate(t *testing.T) {
	_, err := userRepo.Create(context.Background(), domain.User{
		Email:    "test@mail.com",
		Password: "secret",
		Permissions: []domain.PermissionType{
			"can_create",
		},
	})

	if err != nil {
		t.Error(err)
	}

}
