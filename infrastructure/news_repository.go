package infrastructure

import (
	"log"
	"os"

	"github.com/gocolly/colly"

	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type NewsRepository struct {
	c *colly.Collector
}

func NewNewsRepository(c *colly.Collector) repository.NewsRepository {
	return &NewsRepository{c: c}
}

func (r *NewsRepository) Get() ([]*domain.News, error) {
	err := r.c.Post(os.Getenv("APP_DOMAIN")+os.Getenv("APP_LOGIN"),
		map[string]string{
			"c":        "login_2",
			"flg_auto": "1",
			"token_a":  "",
			"id":       os.Getenv("TEST_ID"),
			"pw":       os.Getenv("TEST_PW"),
		})
	if err != nil {
		log.Fatal(err)
	}

	r.c.OnHTML("ul.news_list01 li", func(e *colly.HTMLElement) {
		// 表示
		log.Println("text: ", e.Text)
	})

	r.c.Visit(os.Getenv("APP_DOMAIN") + os.Getenv("APP_NEWS"))

	return []*domain.News{}, nil
}
