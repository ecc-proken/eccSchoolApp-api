package repository

import "github.com/yumekiti/eccSchoolApp-api/domain"

type SigninRepository interface {
	Get(*domain.User) (*domain.Signin, error)
}
