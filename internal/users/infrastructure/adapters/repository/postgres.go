package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yavurb/mobility-payments/internal/common/storage/postgres"
	"github.com/yavurb/mobility-payments/internal/users/domain"
)

type UserRepository struct {
	db *postgres.Queries
}

// TODO: Implement the UnitOfWork pattern
func NewUserRepository(connpool *pgxpool.Pool) domain.Repository {
	db := postgres.New(connpool)

	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user *domain.UserCreate) (*domain.User, error) {
	row, err := r.db.Save(ctx, postgres.SaveParams{
		PublicID: user.PublicID,
		Type:     postgres.UserType(user.Type),
		UserName: user.Name,
		Email:    user.Email,
		Password: user.Password,
		Balance:  user.Balance,
	})
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        row.ID,
		Type:      user.Type,
		PublicID:  user.PublicID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Balance:   user.Balance,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	row, err := r.db.GetByEmail(ctx, email)
	if err != nil {
		log.Printf("Error getting User. Got: %v", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, err
	}

	return &domain.User{
		ID:        row.ID,
		PublicID:  row.PublicID,
		Type:      domain.UserType(row.Type),
		Name:      row.UserName,
		Email:     row.Email,
		Password:  row.Password,
		Balance:   row.Balance,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) GetByPublicID(ctx context.Context, id string) (*domain.User, error) {
	row, err := r.db.GetByPublicID(ctx, id)
	if err != nil {
		log.Printf("Error getting User. Got: %v", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, err
	}

	return &domain.User{
		ID:        row.ID,
		PublicID:  row.PublicID,
		Type:      domain.UserType(row.Type),
		Name:      row.UserName,
		Email:     row.Email,
		Password:  row.Password,
		Balance:   row.Balance,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}, nil
}
