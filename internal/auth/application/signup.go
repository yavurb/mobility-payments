package application

import (
	"context"
	"time"

	"github.com/yavurb/mobility-payments/internal/auth/domain"
	userDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

func (uc *authUsecase) SignUp(ctx context.Context, user_type userDomain.UserType, name, email, password string) (string, string, error) {
	hashedPassword, err := uc.passwordHasher.Hash(password)
	if err != nil {
		return "", "", err
	}

	user, err := uc.userUsecase.Create(ctx, user_type, name, email, hashedPassword)
	if err != nil {
		return "", "", err
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
