package application

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	commonMocks "github.com/yavurb/mobility-payments/internal/common/mocks"
	"github.com/yavurb/mobility-payments/internal/payments/application/mocks"
	"github.com/yavurb/mobility-payments/internal/payments/domain"
	usersDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

func TestVerify(t *testing.T) {
	t.Run("it should verify a transaction by its public ID", func(t *testing.T) {
		want := &domain.Transaction{
			ID:               1,
			Status:           domain.Succeeded,
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
					Status:           domain.RequiresConfirmation,
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
			UpdateTransactionFn: func(ctx context.Context, transactionID string, transaction domain.TransactionUpdate) (*domain.Transaction, error) {
				return &domain.Transaction{
					ID:               1,
					PublicID:         transactionID,
					Status:           transaction.Status,
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
					Type:      usersDomain.Merchant,
					PublicID:  "us_omjnu4m8lsir",
					Name:      "testuser",
					Email:     "testuser@testmail.com",
					Password:  "hashedpassword",
					Balance:   10000, // 100.00 USD
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}, nil
			},
		}

		uc := NewPaymentsUsecase(repository, userUsecase)
		ctx := context.Background()

		got, err := uc.Verify(ctx, want.PublicID, want.Status, want.ReceiverPublicID)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !got.Equal(*want) {
			t.Errorf("Mismatch verifying transaction. (-want,+got):\n%s", cmp.Diff(want, got))
		}
	})
}
