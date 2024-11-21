package application

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/yavurb/mobility-payments/internal/users/application/mocks"
	"github.com/yavurb/mobility-payments/internal/users/domain"
)

func TestUpdateBalance(t *testing.T) {
	t.Run("it should update the balance", func(t *testing.T) {
		var want int64 = 15000

		repository := &mocks.RepositoryMock{
			GetByPublicIDFn: func(ctx context.Context, id string) (*domain.User, error) {
				return &domain.User{
					ID:        1,
					PublicID:  id,
					Name:      "roy",
					Email:     "myemail@tesmail.com",
					Type:      domain.Customer,
					Balance:   100_000,
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}, nil
			},
			UpdateBalanceFn: func(ctx context.Context, id string, amount int64) (int64, error) {
				return amount, nil
			},
		}

		uc := NewUserUsecase(repository)
		ctx := context.Background()

		got, err := uc.UpdateUserBalance(ctx, "us_sj38jdir83", 15000)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if got != want {
			t.Errorf("Mismatch updating balance. (-want,+got):\n%s", cmp.Diff(want, got))
		}
	})
}
