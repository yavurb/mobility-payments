package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yavurb/mobility-payments/internal/common/storage/postgres"
	"github.com/yavurb/mobility-payments/internal/payments/domain"
)

type paymentsRepository struct {
	db *postgres.Queries
}

// TODO: Implement the UnitOfWork pattern
func NewPaymentsRepository(connpool *pgxpool.Pool) domain.Repository {
	db := postgres.New(connpool)

	return &paymentsRepository{
		db: db,
	}
}

func (r *paymentsRepository) CreateTransaction(ctx context.Context, transaction domain.TransactionCreate) (*domain.Transaction, error) {
	row, err := r.db.CreateTransaction(ctx, postgres.CreateTransactionParams{
		Status:      postgres.TransactionStatus(transaction.Status),
		PublicID:    transaction.PublicID,
		Amount:      transaction.Amount,
		Method:      postgres.TransactionPaymentMethod(transaction.Method),
		Description: transaction.Description,
		SenderID:    transaction.SenderID,
		ReceiverID:  transaction.ReceiverID,
	})
	if err != nil {
		return nil, err
	}

	row_, err := r.GetTransaction(ctx, transaction.PublicID)
	if err != nil {
		return nil, err
	}

	return &domain.Transaction{
		ID:               row.ID,
		Status:           transaction.Status,
		PublicID:         transaction.PublicID,
		Amount:           transaction.Amount,
		Method:           transaction.Method,
		Description:      transaction.Description,
		SenderID:         transaction.SenderID,
		ReceiverID:       transaction.ReceiverID,
		ReceiverPublicID: row_.ReceiverPublicID,
		SenderPublicID:   row_.SenderPublicID,
		CreatedAt:        row.CreatedAt.Time,
		UpdatedAt:        row.UpdatedAt.Time,
	}, nil
}

func (r *paymentsRepository) GetTransaction(ctx context.Context, id string) (*domain.Transaction, error) {
	row, err := r.db.GetTransaction(ctx, id)
	if err != nil {
		log.Printf("Error getting Transaction. Got: %v", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrTransactionNotFound
		}

		return nil, err
	}

	return &domain.Transaction{
		ID:               row.ID,
		Status:           domain.TransactionStatus(row.Status),
		PublicID:         row.PublicID,
		Amount:           row.Amount,
		Method:           domain.TransactionMethod(row.Method),
		Description:      row.Description,
		SenderID:         row.SenderID,
		ReceiverID:       row.ReceiverID,
		ReceiverPublicID: row.ReceiverPublicID,
		SenderPublicID:   row.SenderPublicID,
		CreatedAt:        row.CreatedAt.Time,
		UpdatedAt:        row.UpdatedAt.Time,
	}, nil
}

func (r *paymentsRepository) GetSenderTransactions(ctx context.Context, id string) ([]*domain.Transaction, error) {
	rows, err := r.db.GetSenderTransactions(ctx, id)
	if err != nil {
		log.Printf("Error getting Sender Transactions. Got: %v", err)
		return []*domain.Transaction{}, err
	}

	transactions := []*domain.Transaction{}
	for _, row := range rows {
		transactions = append(transactions, &domain.Transaction{
			ID:               row.ID,
			Status:           domain.TransactionStatus(row.Status),
			PublicID:         row.PublicID,
			Amount:           row.Amount,
			Method:           domain.TransactionMethod(row.Method),
			Description:      row.Description,
			SenderID:         row.SenderID,
			ReceiverID:       row.ReceiverID,
			ReceiverPublicID: row.ReceiverPublicID,
			SenderPublicID:   row.SenderPublicID,
			CreatedAt:        row.CreatedAt.Time,
			UpdatedAt:        row.UpdatedAt.Time,
		})
	}

	return transactions, nil
}

func (r *paymentsRepository) GetReceiverTransactions(ctx context.Context, id string) ([]*domain.Transaction, error) {
	rows, err := r.db.GetReceiverTransactions(ctx, id)
	if err != nil {
		return []*domain.Transaction{}, err
	}

	transactions := []*domain.Transaction{}
	for _, row := range rows {
		transactions = append(transactions, &domain.Transaction{
			ID:               row.ID,
			Status:           domain.TransactionStatus(row.Status),
			PublicID:         row.PublicID,
			Amount:           row.Amount,
			Method:           domain.TransactionMethod(row.Method),
			Description:      row.Description,
			SenderID:         row.SenderID,
			ReceiverID:       row.ReceiverID,
			ReceiverPublicID: row.ReceiverPublicID,
			SenderPublicID:   row.SenderPublicID,
			CreatedAt:        row.CreatedAt.Time,
			UpdatedAt:        row.UpdatedAt.Time,
		})
	}

	return transactions, nil
}

func (r *paymentsRepository) UpdateTransaction(ctx context.Context, transactionID string, transaction domain.TransactionUpdate) (*domain.Transaction, error) {
	row, err := r.db.UpdateTransaction(ctx, postgres.UpdateTransactionParams{
		Status:   postgres.TransactionStatus(transaction.Status),
		PublicID: transactionID,
	})
	if err != nil {
		return nil, err
	}

	row_, err := r.GetTransaction(ctx, row.PublicID)
	if err != nil {
		return nil, err
	}

	return &domain.Transaction{
		ID:               row.ID,
		Status:           domain.TransactionStatus(row.Status),
		PublicID:         row.PublicID,
		Amount:           row.Amount,
		Method:           domain.TransactionMethod(row.Method),
		Description:      row.Description,
		SenderID:         row.SenderID,
		ReceiverID:       row.ReceiverID,
		ReceiverPublicID: row_.ReceiverPublicID,
		SenderPublicID:   row_.SenderPublicID,
		CreatedAt:        row.CreatedAt.Time,
		UpdatedAt:        row.UpdatedAt.Time,
	}, nil
}
