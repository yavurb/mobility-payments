package application

import (
	"context"
	"testing"
	"time"

	"github.com/yavurb/mobility-payments/internal/auth/application/mocks"
	"github.com/yavurb/mobility-payments/internal/auth/domain"
	commonMocks "github.com/yavurb/mobility-payments/internal/common/mocks"
	userDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

func TestSignUp(t *testing.T) {
	t.Run("it should sign up a user", func(t *testing.T) {
		wantToken := "somejwttoken"
		wantID := "us_omjnu4m8lsir"

		userUsecase := &commonMocks.UserUsecaseMock{
			CreateFn: func(ctx context.Context, user_type userDomain.UserType, name, email, password string) (*userDomain.User, error) {
				return &userDomain.User{
					ID:        1,
					PublicID:  "us_omjnu4m8lsir",
					Name:      name,
					Email:     email,
					Password:  password,
					Type:      user_type,
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}, nil
			},
		}
		passwordHasher := &mocks.PasswordHasherMock{
			HashFn: func(password string) (string, error) {
				return "hashedpassword", nil
			},
		}
		tokenManager := &mocks.TokenManagerMock{
			GenerateFn: func(payload domain.TokenPayload, duration time.Duration) (string, error) {
				return "somejwttoken", nil
			},
		}

		uc := NewAuthUsecase(userUsecase, passwordHasher, tokenManager)
		ctx := context.Background()

		gotToken, gotID, err := uc.SignUp(ctx, userDomain.Customer, "roy", "roy@email.com", "myunsecuredpassword")
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
