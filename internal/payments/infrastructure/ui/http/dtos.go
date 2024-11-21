package http

import (
	"time"

	"github.com/yavurb/mobility-payments/internal/payments/domain"
)

type PaymentData struct {
	Merchant    string                   `json:"merchant" validate:"required"`
	Method      domain.TransactionMethod `json:"method" validate:"required,oneof=credit_card debit_card bank_transfer"`
	Description string                   `json:"description"`
	Amount      int64                    `json:"amount" validate:"required"`
}

type Transaction struct {
	CreatedAt   time.Time                `json:"created_at"`
	ID          string                   `json:"id"`
	ReceiverID  string                   `json:"receiver_id"`
	SenderID    string                   `json:"sender_id"`
	Description string                   `json:"description"`
	Status      domain.TransactionStatus `json:"status"`
	Method      domain.TransactionMethod `json:"method"`
	Amount      int64                    `json:"amount"`
}

type TransactionList struct {
	Data []Transaction `json:"data"`
}

type TransactionCreated struct {
	ID     string                   `json:"id"`
	Status domain.TransactionStatus `json:"status"`
}

type TransactionVeriyfied struct {
	ID     string                   `json:"id"`
	Status domain.TransactionStatus `json:"status"`
}

type GetTransactionParam struct {
	ID string `param:"id" validate:"required"`
}

type ConfirmationType string

const (
	DeclineConfirmation ConfirmationType = "decline"
	ConfirmConfirmation ConfirmationType = "confirm"
)

type VerifyTransactionData struct {
	ID           string           `param:"id" validate:"required"`
	Confirmation ConfirmationType `json:"confirmation" validate:"required,oneof=decline confirm"`
}
