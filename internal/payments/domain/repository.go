package domain

import "context"

type Repository interface {
	CreateTransaction(ctx context.Context, transaction TransactionCreate) (*Transaction, error)
	GetSenderTransactions(ctx context.Context, senderID string) ([]*Transaction, error)
	GetReceiverTransactions(ctx context.Context, receiverID string) ([]*Transaction, error)
	GetTransaction(ctx context.Context, transactionID string) (*Transaction, error)
	UpdateTransaction(ctx context.Context, transactionID string, transaction TransactionUpdate) (*Transaction, error)
}
