package ports

import (
	"context"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, in domain.User) (domain.User, error)
	FindById(ctx context.Context, id string) (domain.User, error)
}
