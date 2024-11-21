package application

import (
	"context"
)

func (uc *UserUsecase) UpdateUserBalance(ctx context.Context, id string, amount int64) (int64, error) {
	user, err := uc.repository.GetByPublicID(ctx, id)
	if err != nil {
		return 0, err
	}

	newAmount, err := uc.repository.UpdateBalance(ctx, user.PublicID, amount)
	if err != nil {
		return 0, err
	}

	return newAmount, nil
}
