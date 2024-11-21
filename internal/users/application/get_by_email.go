package application

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/users/domain"
)

func (uc *UserUsecase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := uc.repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
