package domain

import (
	"reflect"
	"time"
)

type UserType string

const (
	Customer UserType = "customer"
	Merchant UserType = "merchant"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Type      UserType
	Password  string
	PublicID  string

	Balance int64
	ID      int64
}

type UserCreate struct {
	Type     UserType
	Name     string
	Email    string
	PublicID string
	Password string
	Balance  int64
}

func (u User) Equal(u2 User) bool {
	createdAtDuration := u.CreatedAt.Sub(u2.CreatedAt)
	updatedAtDuration := u.UpdatedAt.Sub(u2.UpdatedAt)

	threshold := time.Second * 5

	if createdAtDuration > threshold || updatedAtDuration > threshold {
		return false
	}

	u.CreatedAt = u2.CreatedAt
	u.UpdatedAt = u2.UpdatedAt

	return reflect.DeepEqual(u, u2)
}

func (u *User) CanDebit(amount int64) error {
	if u.Balance < amount {
		return ErrInsufficientBalance
	}

	return nil
}

func (u *User) Debit(amount int64) error {
	if err := u.CanDebit(amount); err != nil {
		return err
	}

	u.Balance -= amount

	return nil
}

func (u *User) Credit(amount int64) {
	u.Balance += amount
}

func (u UserCreate) CalculateBaseBalance() int64 {
	if u.Type == Customer {
		return 100_000 // 1000USD for Customers
	}

	return 10_000 // 100USD for Merchants
}
