package application

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/users/domain"
)

func (uc *UserUsecase) GetByPublicID(ctx context.Context, id string) (*domain.User, error) {
	user, err := uc.repository.GetByPublicID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
