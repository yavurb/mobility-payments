package mocks

import (
	"time"

	"github.com/yavurb/mobility-payments/internal/auth/domain"
)

type TokenManagerMock struct {
	GenerateFn func(payload domain.TokenPayload, duration time.Duration) (string, error)
	VerifyFn   func(token string) (domain.TokenPayload, error)
}

func (tm *TokenManagerMock) Generate(payload domain.TokenPayload, duration time.Duration) (string, error) {
	return tm.GenerateFn(payload, duration)
}

func (tm *TokenManagerMock) Verify(token string) (domain.TokenPayload, error) {
	return tm.VerifyFn(token)
}
