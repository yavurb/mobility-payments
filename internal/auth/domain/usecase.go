package domain

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/users/domain"
)

type Usecase interface {
	SignUp(context context.Context, user_type domain.UserType, name, email, password string) (string, string, error)
	SignIn(context context.Context, email, password string) (string, string, error)
}
