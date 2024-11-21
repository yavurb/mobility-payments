package domain

import (
	"context"
)

type Usecase interface {
	GetTransactions(ctx context.Context, userID string) ([]*Transaction, error)
	GetTransaction(ctx context.Context, trasactionID, userID string) (*Transaction, error)
	Pay(ctx context.Context, senderID, receiverID, description string, method TransactionMethod, amount int64) (*Transaction, error)
	Verify(ctx context.Context, transactionID string, status TransactionStatus, userID string) (*Transaction, error)
}
