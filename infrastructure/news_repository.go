package infrastructure

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type NewsRepository struct{}

func NewNewsRepository() repository.NewsRepository {
	return &NewsRepository{}
}

func (r *NewsRepository) Get() (*domain.News, error) {
	news := &domain.News{}

	return news, nil
}
