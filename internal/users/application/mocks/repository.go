package mocks

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/users/domain"
)

type RepositoryMock struct {
	SaveFn       func(ctx context.Context, user *domain.UserCreate) (*domain.User, error)
	GetByEmailFn func(ctx context.Context, email string) (*domain.User, error)
}

func (r *RepositoryMock) Save(ctx context.Context, user *domain.UserCreate) (*domain.User, error) {
	return r.SaveFn(ctx, user)
}

func (r *RepositoryMock) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return r.GetByEmailFn(ctx, email)
}
