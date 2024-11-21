package application

import (
	"github.com/yavurb/mobility-payments/internal/payments/domain"
	userDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

const module_prefix string = "tr"

type paymentsUsecase struct {
	paymentsRepository domain.Repository
	userUsecase        userDomain.Usecase
}

func NewPaymentsUsecase(paymentsRepository domain.Repository, userUsecase userDomain.Usecase) domain.Usecase {
	return &paymentsUsecase{
		paymentsRepository: paymentsRepository,
		userUsecase:        userUsecase,
	}
}
