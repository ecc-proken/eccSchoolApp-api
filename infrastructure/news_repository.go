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
	// return []*domain.News{
	// 	&domain.News{
	// 		ID:    1,
	// 		Title: "title1",
	// 		Date:  "date1",
	// 		Link:  "link1",
	// 	},
	// 	&domain.News{
	// 		ID:    2,
	// 		Title: "title2",
	// 		Date:  "date2",
	// 		Link:  "link2",
	// 	},
	// }, nil
	news := []*domain.News{
		&domain.News{
			ID:    1,
			Title: "title1",
			Date:  "date1",
			Link:  "link1",
		},
	}

	return news, nil
}
