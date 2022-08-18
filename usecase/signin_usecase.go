package usecase

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type SigninUsecase interface {
	Get(*domain.User) (*domain.Signin, error)
}

type signinUsecase struct {
	signinRepository repository.SigninRepository
}

func NewSigninUsecase(signinRepository repository.SigninRepository) SigninUsecase {
	return &signinUsecase{signinRepository: signinRepository}
}

func (u *signinUsecase) Get(user *domain.User) (*domain.Signin, error) {
	getSignin, err := u.signinRepository.Get(user)
	if err != nil {
		return nil, err
	}

	return getSignin, nil
}
