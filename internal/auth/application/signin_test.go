package application

import (
	"context"
	"testing"
	"time"

	"github.com/yavurb/mobility-payments/internal/auth/application/mocks"
	"github.com/yavurb/mobility-payments/internal/auth/domain"
	userDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

const (
	wantToken string = "somejwttoken"
	wantID    string = "us_omjnu4m8lsir"
)

func TestSignIn(t *testing.T) {
	t.Run("it should sign up a user", func(t *testing.T) {
		userUsecase := &mocks.UserUsecaseMock{
			GetByEmailFn: func(ctx context.Context, email string) (*userDomain.User, error) {
				return &userDomain.User{
					ID:        1,
					PublicID:  "us_omjnu4m8lsir",
					Name:      "roy",
					Email:     "roy@email.gom",
					Password:  "hashedpassword",
					Type:      userDomain.Customer,
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}, nil
			},
		}
		passwordHasher := &mocks.PasswordHasherMock{
			VerifyFn: func(password, hashedpassword string) (bool, error) {
				return true, nil
			},
		}
		tokenManager := &mocks.TokenManagerMock{
			GenerateFn: func(payload domain.TokenPayload, duration time.Duration) (string, error) {
				return wantToken, nil
			},
		}

		uc := NewAuthUsecase(userUsecase, passwordHasher, tokenManager)
		ctx := context.Background()

		gotToken, gotID, err := uc.SignIn(ctx, "roy@email.com", "myunsecuredpassword")
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if gotToken != wantToken {
			t.Errorf("Mismatch signing up user. Want token: %s. Got token: %s", wantToken, gotToken)
		}

		if gotID != wantID {
			t.Errorf("Mismatch signing up user. Want id: %s. Got id: %s", wantID, gotID)
		}
	})
}
