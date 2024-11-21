package application

import (
	"context"
	"fmt"

	"github.com/yavurb/mobility-payments/internal/payments/domain"
)

func (uc *paymentsUsecase) Verify(ctx context.Context, transactionID string, status domain.TransactionStatus, userID string) (*domain.Transaction, error) {
	receiver, err := uc.userUsecase.GetByPublicID(ctx, userID)
	if err != nil {
		fmt.Println("Error getting recevier at verify: ", err)
		// TODO: Handle not found errors
		return nil, err
	}

	transaction, err := uc.paymentsRepository.GetTransaction(ctx, transactionID)
	if err != nil {
		fmt.Println("Error getting transaction at verify: ", err)
		return nil, err
	}

	if transaction.ReceiverPublicID != userID {
		return nil, domain.ErrTransactionNotFromUser
	}

	if status == domain.Declined {
		sender, err := uc.userUsecase.GetByPublicID(ctx, transaction.SenderPublicID)
		if err != nil {
			// TODO: Handle not found errors
			return nil, err
		}

		err = transaction.Revert(sender, receiver)
		if err != nil {
			return nil, err
		}

		if _, err := uc.userUsecase.UpdateUserBalance(ctx, sender.PublicID, sender.Balance); err != nil {
			return nil, err
		}

		if _, err := uc.userUsecase.UpdateUserBalance(ctx, receiver.PublicID, receiver.Balance); err != nil {
			return nil, err
		}
	}

	transaction_updated, err := uc.paymentsRepository.UpdateTransaction(ctx, transaction.PublicID, domain.TransactionUpdate{
		Status: status,
	})
	if err != nil {
		fmt.Println("Error updating transaction at verify: ", err)
		return nil, err
	}

	return transaction_updated, err
}
