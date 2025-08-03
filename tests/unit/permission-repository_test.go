package unit

import (
	"context"
	"testing"

	"github.com/brandoyts/permission-based-access-control/internal/adapters/persistence/memory"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	err := permissionRepo.Create(context.Background(), "can_create")
	if err != nil {
		t.Error(err)
	}
}

func TestErrCreate(t *testing.T) {
	t.Parallel()

	err := permissionRepo.Create(context.Background(), "can_create")
	if err == nil {
		t.Error("expecting", memory.ErrRecordAlreadyExist)
	}
}
