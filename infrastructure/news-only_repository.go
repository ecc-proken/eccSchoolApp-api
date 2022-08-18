package infrastructure

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type NewsOnlyRepository struct{}

func NewNewsOnlyRepository() repository.NewsOnlyRepository {
	return &NewsOnlyRepository{}
}

func (r *NewsOnlyRepository) Get(user *domain.User) (*domain.NewsOnly, error) {
	return &domain.NewsOnly{}, nil
}
