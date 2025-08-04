package services

import (
	"context"
	"time"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
	"github.com/brandoyts/permission-based-access-control/internal/core/ports"
)

type PermissionService struct {
	repository ports.PermissionRepository
}

func NewPermissionService(repo ports.PermissionRepository) *PermissionService {
	return &PermissionService{repository: repo}
}

func (s *PermissionService) CreatePermission(permission domain.PermissionType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := s.repository.Create(ctx, permission)
	if err != nil {
		return err
	}

	return nil
}
