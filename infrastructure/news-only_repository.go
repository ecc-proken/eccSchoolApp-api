package infrastructure

import (
	"os"

	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type NewsOnlyRepository struct{}

func NewNewsOnlyRepository() repository.NewsOnlyRepository {
	return &NewsOnlyRepository{}
}

func (r *NewsOnlyRepository) Get(id string, user *domain.User) (*domain.NewsOnly, error) {
	c := config.ECCLogin(user)

	c.Visit(os.Getenv("APP_DOMAIN") + os.Getenv("APP_NEWS"))

	return &domain.NewsOnly{
		Title: id,
	}, nil
}
