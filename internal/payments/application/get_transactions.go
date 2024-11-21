package application

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/payments/domain"
	usersDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

func (uc *paymentsUsecase) GetTransactions(ctx context.Context, userID string) ([]*domain.Transaction, error) {
	user, err := uc.userUsecase.GetByPublicID(ctx, userID)
	if err != nil {
		// TODO: handle errors
		return nil, err
	}

	var transactions []*domain.Transaction
	if user.Type == usersDomain.Customer {
		transactions, err = uc.paymentsRepository.GetSenderTransactions(ctx, user.PublicID)
		if err != nil {
			return nil, err
		}
	} else {
		transactions, err = uc.paymentsRepository.GetReceiverTransactions(ctx, user.PublicID)
		if err != nil {
			return nil, err
		}
	}

	return transactions, nil
}
