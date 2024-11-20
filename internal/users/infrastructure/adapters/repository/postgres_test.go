package repository

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yavurb/mobility-payments/internal/users/domain"
	"github.com/yavurb/mobility-payments/testhelpers"
)

func TestSave(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := testhelpers.CreatePostgresContainer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := pgxpool.New(ctx, pgContainer.ConnString)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { conn.Close() })

	repo := NewUserRepository(conn)

	t.Run("it should save a user", func(t *testing.T) {
		testhelpers.CleanDatabase(t, ctx, pgContainer.ConnString)

		want := &domain.User{
			ID:        1,
			Type:      domain.Customer,
			PublicID:  "us_omjnu4m8lsir",
			Name:      "testuser",
			Email:     "testuser@testmail.com",
			Password:  "hashedpassword",
			Balance:   100000, // 1000.00 USD
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		got, err := repo.Save(ctx, &domain.UserCreate{
			Type:     domain.Customer,
			PublicID: "us_omjnu4m8lsir",
			Name:     "testuser",
			Email:    "testuser@testmail.com",
			Password: "hashedpassword",
			Balance:  100000, // 1000.00 USD
		})
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !got.Equal(*want) {
			t.Errorf("Mismatch saving user. (-want,+got):\n%s", cmp.Diff(want, got))
		}
	})
}

func TestGetByEmail(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := testhelpers.CreatePostgresContainer(t, ctx)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := pgxpool.New(ctx, pgContainer.ConnString)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { conn.Close() })

	repo := NewUserRepository(conn)

	t.Run("it should return a user", func(t *testing.T) {
		testhelpers.CleanDatabase(t, ctx, pgContainer.ConnString)

		want := &domain.User{
			ID:        1,
			Type:      domain.Customer,
			PublicID:  "us_omjnu4m8lsir",
			Name:      "testuser",
			Email:     "testuser@testmail.com",
			Password:  "hashedpassword",
			Balance:   100000, // 1000.00 USD
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		_, err := repo.Save(ctx, &domain.UserCreate{
			Type:     domain.Customer,
			PublicID: "us_omjnu4m8lsir",
			Name:     "testuser",
			Email:    "testuser@testmail.com",
			Password: "hashedpassword",
			Balance:  100000, // 1000.00 USD
		})
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		got, err := repo.GetByEmail(ctx, "testuser@testmail.com")
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !got.Equal(*want) {
			t.Errorf("Mismatch saving user. (-want,+got):\n%s", cmp.Diff(want, got))
		}
	})
}
