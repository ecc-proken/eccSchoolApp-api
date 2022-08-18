package repository

import "github.com/yumekiti/eccSchoolApp-api/domain"

type NewsOnlyRepository interface {
	Get(*domain.User) (*domain.NewsOnly, error)
}
