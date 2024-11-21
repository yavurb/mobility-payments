package application

import (
	"github.com/yavurb/mobility-payments/internal/auth/domain"
	userDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

type authUsecase struct {
	userUsecase    userDomain.Usecase
	passwordHasher domain.PasswordHasher
	tokenManager   domain.TokenManager
}

func NewAuthUsecase(userUsecase userDomain.Usecase, passwordHasher domain.PasswordHasher, tokenManager domain.TokenManager) domain.Usecase {
	return &authUsecase{
		userUsecase:    userUsecase,
		passwordHasher: passwordHasher,
		tokenManager:   tokenManager,
	}
}
