package application

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	commonMocks "github.com/yavurb/mobility-payments/internal/common/mocks"
	"github.com/yavurb/mobility-payments/internal/payments/application/mocks"
	"github.com/yavurb/mobility-payments/internal/payments/domain"
	usersDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

func TestPay(t *testing.T) {
	t.Run("it should create a user", func(t *testing.T) {
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

		NewPublicID = func(prefix string) (string, error) {
			return "tr_omjnu4m8lsir", nil
		}
		repository := &mocks.RepositoryMock{
			CreateTransactionFn: func(ctx context.Context, transaction domain.TransactionCreate) (*domain.Transaction, error) {
				return &domain.Transaction{
					ID:               1,
					Status:           transaction.Status,
					PublicID:         transaction.PublicID,
					Amount:           transaction.Amount,
					Method:           transaction.Method,
					Description:      transaction.Description,
					SenderID:         transaction.SenderID,
					ReceiverID:       transaction.ReceiverID,
					SenderPublicID:   want.SenderPublicID,
					ReceiverPublicID: want.ReceiverPublicID,
					CreatedAt:        time.Now().UTC(),
					UpdatedAt:        time.Now().UTC(),
				}, nil
			},
		}
		userUsecase := &commonMocks.UserUsecaseMock{
			GetByPublicIDFn: func(ctx context.Context, id string) (*usersDomain.User, error) {
				user := &usersDomain.User{
					ID:        2,
					Type:      usersDomain.Customer,
					PublicID:  "us_omjnu4m8lsir",
					Name:      "testuser",
					Email:     "testuser@testmail.com",
					Password:  "hashedpassword",
					Balance:   100000, // 1000.00 USD
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}

				if id == "us_omjnu4m8lhol" {
					user = &usersDomain.User{
						ID:        1,
						Type:      usersDomain.Customer,
						PublicID:  "us_omjnu4m8lhol",
						Name:      "testusermerchant",
						Email:     "testusermerchant@testmail.com",
						Password:  "hashedpassword",
						Balance:   10000, // 100.00 USD
						CreatedAt: time.Now().UTC(),
						UpdatedAt: time.Now().UTC(),
					}
				}

				return user, nil
			},
			UpdateUserBalanceFn: func(ctx context.Context, id string, amount int64) (int64, error) {
				return amount, nil
			},
		}

		uc := NewPaymentsUsecase(repository, userUsecase)
		ctx := context.Background()

		got, err := uc.Pay(ctx, want.SenderPublicID, want.ReceiverPublicID, want.Description, want.Method, want.Amount)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		rgx := regexp.MustCompile(`tr_[a-zA-Z0-9]{5}`)
		if !rgx.MatchString(got.PublicID) {
			t.Errorf("Expected PublicID to match the regex %s, got: %s", rgx.String(), got.PublicID)
		}

		if !got.Equal(*want) {
			t.Errorf("Mismatch creating user. (-want,+got):\n%s", cmp.Diff(want, got))
		}
	})
}
