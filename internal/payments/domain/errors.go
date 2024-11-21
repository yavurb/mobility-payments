package domain

import "errors"

var (
	ErrTransactionNotFound    = errors.New("transaction not found")
	ErrTransactionNotFromUser = errors.New("transaction not from user")
	ErrInsufficientBalance    = errors.New("sender user has insufficient balance to process the transaction")
)
