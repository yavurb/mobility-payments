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

func (u UserCreate) CalculateBaseBalance() int64 {
	if u.Type == Customer {
		return 100_000 // 1000USD for Customers
	}

	return 10_000 // 100USD for Merchants
}
