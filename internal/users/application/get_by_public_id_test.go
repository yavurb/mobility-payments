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

func TestGetByPublicID(t *testing.T) {
	t.Run("it should get a user by its public ID", func(t *testing.T) {
		want := &domain.User{
			ID:        1,
			Type:      domain.Customer,
			PublicID:  "us_omjnu4m8lsir",
			Name:      "testuser",
			Email:     "testuser@testmail.com",
			Balance:   100_000,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		repository := &mocks.RepositoryMock{
			GetByPublicIDFn: func(ctx context.Context, id string) (*domain.User, error) {
				if id != want.PublicID {
					return nil, domain.ErrUserNotFound
				}

				return &domain.User{
					ID:        1,
					PublicID:  id,
					Name:      want.Name,
					Email:     want.Email,
					Type:      domain.Customer,
					Balance:   100_000,
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}, nil
			},
		}

		uc := NewUserUsecase(repository)
		ctx := context.Background()

		got, err := uc.GetByPublicID(ctx, want.PublicID)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		rgx := regexp.MustCompile(`us_[a-zA-Z0-9]{5}`)
		if !rgx.MatchString(got.PublicID) {
			t.Errorf("Expected PublicID to match the regex %s, got: %s", rgx.String(), got.PublicID)
		}

		if !got.Equal(*want) {
			t.Errorf("Mismatch getting user by public ID. (-want,+got):\n%s", cmp.Diff(want, got))
		}
	})

	t.Run("it should return an error if does not exists", func(t *testing.T) {
		repository := &mocks.RepositoryMock{
			GetByPublicIDFn: func(ctx context.Context, id string) (*domain.User, error) {
				return nil, domain.ErrUserNotFound
			},
		}

		uc := NewUserUsecase(repository)
		ctx := context.Background()

		_, err := uc.GetByPublicID(ctx, "us_noexists")
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if !errors.Is(err, domain.ErrUserNotFound) {
			t.Errorf("Expected ErrUserNotFound. Got %v", err)
		}
	})
}
