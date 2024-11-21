package application

import (
	"context"
	"errors"
	"time"

	"github.com/yavurb/mobility-payments/internal/auth/domain"
	userDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

func (uc *authUsecase) SignIn(ctx context.Context, email, password string) (string, string, error) {
	user, err := uc.userUsecase.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, userDomain.ErrUserNotFound) {
			return "", "", domain.ErrUserNotFound
		}

		return "", "", err
	}

	// TODO: Add a test case
	validPassword, err := uc.passwordHasher.Verify(password, user.Password)
	if err != nil {
		return "", "", err
	}

	if !validPassword {
		return "", "", domain.ErrInvalidCredentials
	}

	token, err := uc.tokenManager.Generate(
		domain.TokenPayload{
			ID:   user.PublicID,
			Type: user.Type,
		},
		time.Hour*24,
	)
	if err != nil {
		return "", "", err
	}

	return token, user.PublicID, nil
}
