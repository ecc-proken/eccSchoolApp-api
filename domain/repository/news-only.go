package repository

import "github.com/yumekiti/eccSchoolApp-api/domain"

type NewsOnlyRepository interface {
	Get(id string, user *domain.User) (*domain.NewsOnly, error)
}
