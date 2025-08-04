package services

import (
	"context"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
	"github.com/brandoyts/permission-based-access-control/internal/core/ports"
)

type UserService struct {
	repository ports.UserRepository
}

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, in domain.User) (domain.User, error) {
	return us.repository.Create(ctx, in)
}

func (us *UserService) FindUserById(ctx context.Context, id string) (domain.User, error) {
	return us.repository.FindById(ctx, id)
}

func (us *UserService) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return us.repository.FindByEmail(ctx, email)
}
