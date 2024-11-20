package application

import (
	"context"
	"errors"

	"github.com/yavurb/mobility-payments/internal/pkg/ids"
	"github.com/yavurb/mobility-payments/internal/users/domain"
)

var NewPublicID = ids.NewPublicID

func (uc *UserUsecase) Create(ctx context.Context, user_type domain.UserType, name, email, password string) (*domain.User, error) {
	user_, err := uc.repository.GetByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, domain.ErrUserNotFound) {
			return nil, err
		}
	} else if user_ != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	publicID, _ := NewPublicID(module_prefix)
	user := &domain.UserCreate{
		PublicID: publicID,
		Type:     user_type,
		Name:     name,
		Email:    email,
		Password: password,
	}

	balance := user.CalculateBaseBalance()
	user.Balance = balance

	return uc.repository.Save(ctx, user)
}
