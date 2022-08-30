package infrastructure

import (
	"os"
	"strings"

	"github.com/gocolly/colly"
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

	// 返す値の初期化
	var title string
	var body string
	var date string
	var tag string
	var attachment []string

	// title
	c.OnHTML(".title", func(e *colly.HTMLElement) {
		title = e.Text
	})
	// body
	c.OnHTML(".news div ~ div", func(e *colly.HTMLElement) {
		html, err := e.DOM.Html()
		if err != nil {
			return
		}
		for body == "" {
			body = html
		}
		body = strings.Replace(body, "\n", "", -1)
	})
	// date
	c.OnHTML(".detail_title01 div", func(e *colly.HTMLElement) {
		date = strings.Split(e.Text, " ")[0]
	})
	// tag
	c.OnHTML(".icon01", func(e *colly.HTMLElement) {
		tag = e.Text
	})
	// attachment
	c.OnHTML(".main article p", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Attr("href") == "../" {
				return
			}
			attachment = append(attachment, e.Attr("href"))
		})
	})

	c.Visit(os.Getenv("APP_DOMAIN") + os.Getenv("APP_NEWS") + os.Getenv("APP_NEWS_ONLY_FRONT") + id + os.Getenv("APP_NEWS_ONLY_BACK"))

	// 返す値から newsOnly を作成
	newsOnly := &domain.NewsOnly{
		Title:      title,
		Body:       body,
		Date:       date,
		Tag:        tag,
		Attachment: attachment,
	}
	return newsOnly, nil
}
