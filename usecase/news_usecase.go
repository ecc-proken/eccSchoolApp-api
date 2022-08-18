package usecase

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type NewsUsecase interface {
	Get(*domain.User) ([]*domain.News, error)
}

type newsUsecase struct {
	newsRepository repository.NewsRepository
}

func NewNewsUsecase(newsRepository repository.NewsRepository) NewsUsecase {
	return &newsUsecase{newsRepository: newsRepository}
}

func (u *newsUsecase) Get(user *domain.User) ([]*domain.News, error) {
	getNews, err := u.newsRepository.Get(user)
	if err != nil {
		return nil, err
	}

	return getNews, nil
}
