package ports

import (
	"context"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
)

type PermissionRepository interface {
	Create(ctx context.Context, in domain.PermissionType) error
	FindOne(ctx context.Context, in domain.PermissionType) (domain.Permission, error)
}
