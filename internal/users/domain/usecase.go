package domain

import "context"

type Usecase interface {
	Create(context context.Context, user_type UserType, name, email, password string) (*User, error)
	GetByEmail(context context.Context, email string) (*User, error)
}
