package domain

import (
	"time"

	"github.com/yavurb/mobility-payments/internal/users/domain"
)

type TokenPayload struct {
	ID   string
	Type domain.UserType
}

type TokenManager interface {
	Generate(payload TokenPayload, duration time.Duration) (string, error)
	Verify(token string) (TokenPayload, error)
}
