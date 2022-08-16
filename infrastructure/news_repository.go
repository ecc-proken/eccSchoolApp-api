package infrastructure

import (
	"log"
	"os"
	"strings"

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

	id := []string{}
	title := []string{}
	date := []string{}
	tag := []string{}
	link := []string{}

	r.c.OnHTML("ul.news_list01 li", func(e *colly.HTMLElement) {
		// id取得
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			id = append(id, strings.Split(strings.Split(e.Attr("href"), "=")[2], "&")[0])
		})
		// title取得
		e.ForEach("dd", func(_ int, e *colly.HTMLElement) {
			title = append(title, e.Text)
		})
		// date取得
		e.ForEach("dt", func(_ int, e *colly.HTMLElement) {
			date = append(date, strings.Split(e.Text, " ")[0])
		})
		//tag取得
		e.ForEach("dt", func(_ int, e *colly.HTMLElement) {
			tag = append(tag, strings.Split(e.Text, " ")[1])
		})
		// link取得
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			link = append(link, os.Getenv("APP_DOMAIN")+os.Getenv("APP_NEWS_LINK")+e.Attr("href")[2:])
		})
	})

	r.c.Visit(os.Getenv("APP_DOMAIN") + os.Getenv("APP_NEWS"))

	// news作成
	news := []*domain.News{}
	for i := 0; i < len(id); i++ {
		news = append(news, &domain.News{
			ID:    id[i],
			Title: title[i],
			Date:  date[i],
			Tag:   tag[i],
			Link:  link[i],
		})
	}
	return news, nil
}
