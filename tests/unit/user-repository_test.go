package unit

import (
	"context"
	"fmt"
	"testing"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
)

func TestUserRepositoryCreate(t *testing.T) {
	user, err := userRepo.Create(context.Background(), domain.User{
		Email:    "test@mail.com",
		Password: "secret",
		Permissions: []domain.PermissionType{
			domain.PermissionCreate,
			domain.PermissionUpdate,
		},
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(user)

}
