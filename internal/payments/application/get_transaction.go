package application

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/payments/domain"
)

func (uc *paymentsUsecase) GetTransaction(ctx context.Context, trasactionID, userID string) (*domain.Transaction, error) {
	user, err := uc.userUsecase.GetByPublicID(ctx, userID)
	if err != nil {
		// TODO: handle errors
		return nil, err
	}

	transaction, err := uc.paymentsRepository.GetTransaction(ctx, trasactionID)
	if err != nil {
		// TODO: handle errors
		return nil, err
	}

	if transaction.ReceiverID != user.ID && transaction.SenderID != user.ID {
		return nil, domain.ErrTransactionNotFromUser
	}

	return transaction, nil
}
