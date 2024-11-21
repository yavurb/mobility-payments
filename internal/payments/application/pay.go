package application

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/payments/domain"
	"github.com/yavurb/mobility-payments/internal/pkg/ids"
)

var NewPublicID = ids.NewPublicID

func (uc *paymentsUsecase) Pay(ctx context.Context, senderID, receiverID, description string, method domain.TransactionMethod, amount int64) (*domain.Transaction, error) {
	receiver, err := uc.userUsecase.GetByPublicID(ctx, receiverID)
	if err != nil {
		// TODO: Handle not found errors
		return nil, err
	}

	sender, err := uc.userUsecase.GetByPublicID(ctx, senderID)
	if err != nil {
		// TODO: Handle not found errors
		return nil, err
	}

	publicID, err := NewPublicID(module_prefix)
	if err != nil {
		return nil, err
	}

	// TODO: Add test case
	transaction_ := domain.TransactionCreate{
		PublicID:    publicID,
		Method:      method,
		Description: description,
		ReceiverID:  receiver.ID,
		SenderID:    sender.ID,
		Amount:      amount,
	}
	transaction_.Status = transaction_.GetInitialStatus()

	err = transaction_.Apply(sender, receiver)
	if err != nil {
		return nil, err
	}

	transaction, err := uc.paymentsRepository.CreateTransaction(ctx, transaction_)
	if err != nil {
		// TODO: Handle not found errors
		return nil, err
	}

	if _, err := uc.userUsecase.UpdateUserBalance(ctx, sender.PublicID, sender.Balance); err != nil {
		return nil, err
	}

	if _, err := uc.userUsecase.UpdateUserBalance(ctx, receiver.PublicID, receiver.Balance); err != nil {
		return nil, err
	}

	return transaction, err
}
