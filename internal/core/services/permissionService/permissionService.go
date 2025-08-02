package permissionService

import (
	"context"
	"time"

	"github.com/brandoyts/permission-based-access-control/internal/core/ports"
)

type Service struct {
	repository ports.PermissionRepository
}

func New(repo ports.PermissionRepository) *Service {
	return &Service{repository: repo}
}

func (s *Service) CreatePermission(permission string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := s.repository.Create(ctx, permission)
	if err != nil {
		return err
	}

	return nil
}
