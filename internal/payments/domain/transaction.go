package domain

import (
	"reflect"
	"time"

	usersDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

type TransactionStatus string

const (
	RequiresConfirmation TransactionStatus = "requires_confirmation"
	Succeeded            TransactionStatus = "succeeded"
	Declined             TransactionStatus = "declined"
)

type TransactionMethod string

const (
	CreditCard   TransactionMethod = "credit_card"
	DebitCard    TransactionMethod = "debit_card"
	BankTransfer TransactionMethod = "bank_transfer"
)

type Transaction struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	PublicID         string
	SenderPublicID   string
	ReceiverPublicID string
	Description      string
	Status           TransactionStatus
	Method           TransactionMethod
	ID               int64
	ReceiverID       int64
	SenderID         int64
	Amount           int64
}

type TransactionCreate struct {
	PublicID    string
	Status      TransactionStatus
	Method      TransactionMethod
	Description string
	ReceiverID  int64
	SenderID    int64
	Amount      int64
}

type TransactionUpdate struct {
	Status TransactionStatus
}

func (t Transaction) Equal(t2 Transaction) bool {
	createdAtDuration := t.CreatedAt.Sub(t2.CreatedAt)
	updatedAtDuration := t.UpdatedAt.Sub(t2.UpdatedAt)

	threshold := time.Second * 5

	if createdAtDuration > threshold || updatedAtDuration > threshold {
		return false
	}

	t.CreatedAt = t2.CreatedAt
	t.UpdatedAt = t2.UpdatedAt

	return reflect.DeepEqual(t, t2)
}

func (t *Transaction) Revert(sender *usersDomain.User, receiver *usersDomain.User) error {
	if err := receiver.Debit(t.Amount); err != nil {
		return ErrInsufficientBalance
	}

	sender.Credit(t.Amount)

	return nil
}

func (tc TransactionCreate) GetInitialStatus() TransactionStatus {
	return RequiresConfirmation
}

func (tc *TransactionCreate) Apply(sender *usersDomain.User, receiver *usersDomain.User) error {
	if err := sender.Debit(tc.Amount); err != nil {
		return ErrInsufficientBalance
	}

	receiver.Credit(tc.Amount)

	return nil
}
