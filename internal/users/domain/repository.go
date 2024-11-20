package domain

import "context"

type Repository interface {
	Save(ctx context.Context, user *UserCreate) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByPublicID(ctx context.Context, id string) (*User, error)
}
