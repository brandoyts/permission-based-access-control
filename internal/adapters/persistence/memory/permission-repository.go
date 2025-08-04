package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
)

const (
	ErrRecordAlreadyExist = "record already exist"
)

type PermissionRepository struct {
	store map[domain.PermissionType]domain.Permission
	mu    sync.RWMutex
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{
		store: map[domain.PermissionType]domain.Permission{},
		mu:    sync.RWMutex{},
	}
}

func (pr *PermissionRepository) Create(ctx context.Context, in domain.PermissionType) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	_, exist := pr.store[in]
	if exist {
		return errors.New(ErrRecordAlreadyExist)
	}

	permission := domain.Permission{
		Type:      in,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pr.store[in] = permission

	return nil
}
