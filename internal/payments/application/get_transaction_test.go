package application

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	commonMocks "github.com/yavurb/mobility-payments/internal/common/mocks"
	"github.com/yavurb/mobility-payments/internal/payments/application/mocks"
	"github.com/yavurb/mobility-payments/internal/payments/domain"
	usersDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

func TestGetTransaction(t *testing.T) {
	t.Run("it should get a transaction by its public ID", func(t *testing.T) {
		want := &domain.Transaction{
			ID:               1,
			Status:           domain.RequiresConfirmation,
			PublicID:         "tr_omjnu4m8lsir",
			Amount:           1000, // 10 USD
			Method:           domain.BankTransfer,
			Description:      "Fish and Fries",
			SenderPublicID:   "us_omjnu4m8lsir",
			ReceiverPublicID: "us_omjnu4m8lhol",
			SenderID:         2,
			ReceiverID:       1,
			CreatedAt:        time.Now().UTC(),
			UpdatedAt:        time.Now().UTC(),
		}

		repository := &mocks.RepositoryMock{
			GetTransactionFn: func(ctx context.Context, id string) (*domain.Transaction, error) {
				return &domain.Transaction{
					ID:               1,
					PublicID:         id,
					Status:           want.Status,
					Amount:           want.Amount,
					Method:           want.Method,
					Description:      want.Description,
					SenderID:         want.SenderID,
					ReceiverID:       want.ReceiverID,
					SenderPublicID:   want.SenderPublicID,
					ReceiverPublicID: want.ReceiverPublicID,
					CreatedAt:        time.Now().UTC(),
					UpdatedAt:        time.Now().UTC(),
				}, nil
			},
		}
		userUsecase := &commonMocks.UserUsecaseMock{
			GetByPublicIDFn: func(ctx context.Context, id string) (*usersDomain.User, error) {
				return &usersDomain.User{
					ID:        1,
					Type:      usersDomain.Customer,
					PublicID:  "us_omjnu4m8lsir",
					Name:      "testuser",
					Email:     "testuser@testmail.com",
					Password:  "hashedpassword",
					Balance:   100000, // 1000.00 USD
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}, nil
			},
		}

		uc := NewPaymentsUsecase(repository, userUsecase)
		ctx := context.Background()

		got, err := uc.GetTransaction(ctx, want.PublicID, want.SenderPublicID)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		rgx := regexp.MustCompile(`tr_[a-zA-Z0-9]{5}`)
		if !rgx.MatchString(got.PublicID) {
			t.Errorf("Expected PublicID to match the regex %s, got: %s", rgx.String(), got.PublicID)
		}

		if !got.Equal(*want) {
			t.Errorf("Mismatch getting user by public ID. (-want,+got):\n%s", cmp.Diff(want, got))
		}
	})

	t.Run("it should return an error if the transaction does not exists", func(t *testing.T) {
		repository := &mocks.RepositoryMock{
			GetTransactionFn: func(ctx context.Context, id string) (*domain.Transaction, error) {
				return nil, domain.ErrTransactionNotFound
			},
		}
		userUsecase := &commonMocks.UserUsecaseMock{
			GetByPublicIDFn: func(ctx context.Context, id string) (*usersDomain.User, error) {
				return &usersDomain.User{
					ID:        1,
					Type:      usersDomain.Customer,
					PublicID:  "us_omjnu4m8lsir",
					Name:      "testuser",
					Email:     "testuser@testmail.com",
					Password:  "hashedpassword",
					Balance:   100000, // 1000.00 USD
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}, nil
			},
		}

		uc := NewPaymentsUsecase(repository, userUsecase)
		ctx := context.Background()

		_, err := uc.GetTransaction(ctx, "tr_noexists", "us_noexists")
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if !errors.Is(err, domain.ErrTransactionNotFound) {
			t.Errorf("Expected ErrUserNotFound. Got %v", err)
		}
	})
}
