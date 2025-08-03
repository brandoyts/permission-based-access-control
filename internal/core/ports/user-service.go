package ports

import (
	"context"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, in domain.User) (domain.User, error)
	FindUserById(ctx context.Context, id string) (domain.User, error)
}
