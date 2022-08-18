package usecase

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type NewsOnlyUsecase interface {
	Get(*domain.User) (*domain.NewsOnly, error)
}

type newsOnlyUsecase struct {
	newsOnlyRepository repository.NewsOnlyRepository
}

func NewNewsOnlyUsecase(newsOnlyRepository repository.NewsOnlyRepository) NewsOnlyUsecase {
	return &newsOnlyUsecase{newsOnlyRepository: newsOnlyRepository}
}

func (u *newsOnlyUsecase) Get(user *domain.User) (*domain.NewsOnly, error) {
	getNewsOnly, err := u.newsOnlyRepository.Get(user)
	if err != nil {
		return nil, err
	}

	return getNewsOnly, nil
}
