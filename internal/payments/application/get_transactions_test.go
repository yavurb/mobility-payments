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

func TestGetTransactions(t *testing.T) {
	want := []*domain.Transaction{
		{
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
		},
		{
			ID:               2,
			Status:           domain.RequiresConfirmation,
			PublicID:         "tr_omjnu4m8lsol",
			Amount:           1589, // 15.89 USD
			Method:           domain.BankTransfer,
			Description:      "Fish and Fries 2",
			SenderPublicID:   "us_omjnu4m8lsir",
			ReceiverPublicID: "us_omjnu4m8lhol",
			SenderID:         2,
			ReceiverID:       1,
			CreatedAt:        time.Now().UTC(),
			UpdatedAt:        time.Now().UTC(),
		},
	}

	repository := &mocks.RepositoryMock{
		GetSenderTransactionsFn: func(ctx context.Context, senderID string) ([]*domain.Transaction, error) {
			return want, nil
		},
		GetReceiverTransactionsFn: func(ctx context.Context, receiverID string) ([]*domain.Transaction, error) {
			return want, nil
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

	transactions, err := uc.GetTransactions(context.Background(), "us_omjnu4m8lsir")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	for i, got := range transactions {
		if !got.Equal(*want[i]) {
			t.Errorf("Mismatch getting transactions. (-want,+got):\n%s", cmp.Diff(want[i], got))
		}
	}
}
