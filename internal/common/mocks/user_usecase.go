package mocks

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/users/domain"
)

type UserUsecaseMock struct {
	CreateFn        func(ctx context.Context, user_type domain.UserType, name, email, password string) (*domain.User, error)
	GetByEmailFn    func(ctx context.Context, email string) (*domain.User, error)
	GetByPublicIDFn func(ctx context.Context, id string) (*domain.User, error)
}

func (uu *UserUsecaseMock) Create(ctx context.Context, user_type domain.UserType, name, email, password string) (*domain.User, error) {
	return uu.CreateFn(ctx, user_type, name, email, password)
}

func (uu *UserUsecaseMock) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return uu.GetByEmailFn(ctx, email)
}

func (uu *UserUsecaseMock) GetByPublicID(ctx context.Context, id string) (*domain.User, error) {
	return uu.GetByPublicIDFn(ctx, id)
}
