package ports

import (
	"context"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
)

type PermissionRepository interface {
	Create(ctx context.Context, in string) error
	FindOne(ctx context.Context, in string) (domain.Permission, error)
}
