package unit

import (
	"context"
	"testing"

	"github.com/brandoyts/permission-based-access-control/internal/infrastructure/persistence/memory/permissionRepository"
)

var repo = permissionRepository.New()

func TestCreate(t *testing.T) {
	t.Parallel()

	err := repo.Create(context.Background(), "can_create")
	if err != nil {
		t.Error(err)
	}
}

func TestErrCreate(t *testing.T) {
	t.Parallel()

	err := repo.Create(context.Background(), "can_creates")
	if err == nil {
		t.Error("expecting", permissionRepository.ErrRecordAlreadyExist)
	}
}
