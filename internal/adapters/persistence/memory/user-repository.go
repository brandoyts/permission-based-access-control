package memory

import (
	"context"
	"sync"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
	"github.com/google/uuid"
)

type UserRepository struct {
	store map[string]domain.User
	mu    sync.RWMutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		store: map[string]domain.User{},
		mu:    sync.RWMutex{},
	}
}

func (ur *UserRepository) Create(ctx context.Context, in domain.User) (domain.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	userId := uuid.NewString()

	user := domain.User{
		ID:          userId,
		Email:       in.Email,
		Password:    in.Password,
		Permissions: in.Permissions,
		// CreatedAt:   time.Now(),
		// UpdatedAt:   time.Now(),
	}

	ur.store[userId] = user

	return user, nil
}

func (ur *UserRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	user, exist := ur.store[id]
	if !exist {
		return domain.User{}, nil
	}

	return user, nil
}
