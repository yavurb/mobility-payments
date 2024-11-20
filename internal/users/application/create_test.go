package application

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/yavurb/mobility-payments/internal/users/application/mocks"
	"github.com/yavurb/mobility-payments/internal/users/domain"
)

func TestCreateUser(t *testing.T) {
	t.Run("it should create a user", func(t *testing.T) {
		want := &domain.User{
			ID:        1,
			Type:      domain.Customer,
			PublicID:  "us_omjnu4m8lsir",
			Name:      "testuser",
			Email:     "testuser@testmail.com",
			Password:  "hashedpassword",
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		NewPublicID = func(prefix string) (string, error) {
			return "us_omjnu4m8lsir", nil
		}
		repository := &mocks.RepositoryMock{
			SaveFn: func(ctx context.Context, user *domain.UserCreate) (*domain.User, error) {
				return &domain.User{
					ID:        1,
					PublicID:  user.PublicID,
					Name:      user.Name,
					Email:     user.Email,
					Password:  user.Password,
					Type:      user.Type,
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}, nil
			},
			GetByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
				return nil, domain.ErrUserNotFound
			},
		}

		uc := NewUserUsecase(repository)
		ctx := context.Background()

		got, err := uc.Create(ctx, want.Type, want.Name, want.Email, want.Password)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		rgx := regexp.MustCompile(`us_[a-zA-Z0-9]{5}`)
		if !rgx.MatchString(got.PublicID) {
			t.Errorf("Expected PublicID to match the regex %s, got: %s", rgx.String(), got.PublicID)
		}

		if !got.Equal(*want) {
			t.Errorf("Mismatch creating user. (-want,+got):\n%s", cmp.Diff(want, got))
		}
	})

	t.Run("it should return an error if user already exists", func(t *testing.T) {
		NewPublicID = func(prefix string) (string, error) {
			return "us_omjnu4m8lsir", nil
		}
		repository := &mocks.RepositoryMock{
			GetByEmailFn: func(ctx context.Context, email string) (*domain.User, error) {
				return nil, domain.ErrUserAlreadyExists
			},
		}

		uc := NewUserUsecase(repository)
		ctx := context.Background()

		_, err := uc.Create(ctx, domain.Customer, "Roy", "email@email.com", "myunsafepassword")
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if !errors.Is(err, domain.ErrUserAlreadyExists) {
			t.Errorf("Expected ErrUserAlreadyExists. Got: %v", err)
		}
	})
}
