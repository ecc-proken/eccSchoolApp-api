package infrastructure

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type SigninRepository struct{}

func NewSigninRepository() repository.SigninRepository {
	return &SigninRepository{}
}

func (r *SigninRepository) Get(user *domain.User) (*domain.Signin, error) {
	return &domain.Signin{}, nil
}
