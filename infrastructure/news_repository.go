package infrastructure

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type NewsRepository struct{}

func NewNewsRepository() repository.NewsRepository {
	return &NewsRepository{}
}

func (r *NewsRepository) Get() ([]*domain.News, error) {
	news := []*domain.News{
		{
			ID:    1,
			Title: "title1",
			Date:  "date1",
			Link:  "link1",
		},
		{
			ID:    2,
			Title: "title1",
			Date:  "date1",
			Link:  "link1",
		},
	}

	return news, nil
}
