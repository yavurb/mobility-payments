package mocks

import (
	"context"

	"github.com/yavurb/mobility-payments/internal/payments/domain"
)

type RepositoryMock struct {
	CreateTransactionFn       func(ctx context.Context, transaction domain.TransactionCreate) (*domain.Transaction, error)
	GetSenderTransactionsFn   func(ctx context.Context, senderID string) ([]*domain.Transaction, error)
	GetReceiverTransactionsFn func(ctx context.Context, receiverID string) ([]*domain.Transaction, error)
	GetTransactionFn          func(ctx context.Context, transactionID string) (*domain.Transaction, error)
	UpdateTransactionFn       func(ctx context.Context, transactionID string, transaction domain.TransactionUpdate) (*domain.Transaction, error)
}

func (r *RepositoryMock) CreateTransaction(ctx context.Context, transaction domain.TransactionCreate) (*domain.Transaction, error) {
	return r.CreateTransactionFn(ctx, transaction)
}

func (r *RepositoryMock) GetSenderTransactions(ctx context.Context, senderID string) ([]*domain.Transaction, error) {
	return r.GetSenderTransactionsFn(ctx, senderID)
}

func (r *RepositoryMock) GetReceiverTransactions(ctx context.Context, receiverID string) ([]*domain.Transaction, error) {
	return r.GetReceiverTransactionsFn(ctx, receiverID)
}

func (r *RepositoryMock) GetTransaction(ctx context.Context, transactionID string) (*domain.Transaction, error) {
	return r.GetTransactionFn(ctx, transactionID)
}

func (r *RepositoryMock) UpdateTransaction(ctx context.Context, transactionID string, transaction domain.TransactionUpdate) (*domain.Transaction, error) {
	return r.UpdateTransactionFn(ctx, transactionID, transaction)
}
