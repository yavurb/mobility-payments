package application

import "github.com/yavurb/mobility-payments/internal/users/domain"

type UserUsecase struct {
	repository domain.Repository
}

func NewUserUsecase(repository domain.Repository) domain.Usecase {
	return &UserUsecase{repository: repository}
}

const module_prefix string = "us"
