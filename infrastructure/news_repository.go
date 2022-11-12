package infrastructure

import (
	"os"
	"strings"

	"github.com/gocolly/colly"

	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type NewsRepository struct{}

func NewNewsRepository() repository.NewsRepository {
	return &NewsRepository{}
}

func (r *NewsRepository) Get(user *domain.User) ([]*domain.News, error) {
	c := config.ECCLogin(user)

	// 返す値の初期化
	var id []string
	var title []string
	var date []string
	var tag []string
	var link []string

	// ニュースを取得し、それぞれを配列に格納
	c.OnHTML("ul.news_list01 li", func(e *colly.HTMLElement) {
		// id
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			id = append(id, strings.Split(strings.Split(e.Attr("href"), "=")[2], "&")[0])
		})
		// title
		e.ForEach("dd", func(_ int, e *colly.HTMLElement) {
			title = append(title, e.Text)
		})
		// date
		e.ForEach("dt", func(_ int, e *colly.HTMLElement) {
			date = append(date, strings.Replace(strings.Split(e.Text, " ")[0], ".", "/", -1))
		})
		//tag
		e.ForEach("dt", func(_ int, e *colly.HTMLElement) {
			tag = append(tag, strings.Join(strings.Split(e.Text, " ")[1:], ""))
		})
		// link
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			link = append(link, os.Getenv("APP_DOMAIN")+"/app/news/"+e.Attr("href")[2:])
		})
	})

	// ニュースのリンク指定
	c.Visit(os.Getenv("APP_DOMAIN") + "/app/news/?c=news")

	// 返す値から news を作成
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
