package ports

import (
	"context"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, in domain.User) error
	FindById(ctx context.Context, id string) (domain.User, error)
	FindOne(ctx context.Context, in domain.User) (domain.User, error)
}
