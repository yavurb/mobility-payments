package mocks

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/users/domain"
)

type RepositoryMock struct {
	SaveFn          func(ctx context.Context, user *domain.UserCreate) (*domain.User, error)
	GetByEmailFn    func(ctx context.Context, email string) (*domain.User, error)
	GetByPublicIDFn func(ctx context.Context, id string) (*domain.User, error)
	UpdateBalanceFn func(ctx context.Context, id string, amount int64) (int64, error)
}

func (r *RepositoryMock) Save(ctx context.Context, user *domain.UserCreate) (*domain.User, error) {
	return r.SaveFn(ctx, user)
}

func (r *RepositoryMock) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return r.GetByEmailFn(ctx, email)
}

func (r *RepositoryMock) GetByPublicID(ctx context.Context, id string) (*domain.User, error) {
	return r.GetByPublicIDFn(ctx, id)
}

func (r *RepositoryMock) UpdateBalance(ctx context.Context, id string, amount int64) (int64, error) {
	return r.UpdateBalanceFn(ctx, id, amount)
}
