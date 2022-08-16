package usecase

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type NewsUsecase interface {
	Get() (*domain.News, error)
}

type newsUsecase struct {
	newsRepository repository.NewsRepository
}

func NewNewsUsecase(newsRepository repository.NewsRepository) NewsUsecase {
	return &newsUsecase{newsRepository: newsRepository}
}

func (u *newsUsecase) Get() (*domain.News, error) {
	getNews, err := u.newsRepository.Get()
	if err != nil {
		return nil, err
	}

	return getNews, nil
}
