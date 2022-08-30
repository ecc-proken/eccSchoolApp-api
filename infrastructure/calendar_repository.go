package infrastructure

import (
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type CalendarRepository struct{}

func NewCalendarRepository() repository.CalendarRepository {
	return &CalendarRepository{}
}

func (r *CalendarRepository) Get(year, month string, user *domain.User) ([]*domain.Calendar, error) {
	c := config.ECCLogin(user)

	// 返す値の初期化
	var day []string
	var title [][]string
	var link [][]string

	c.OnHTML("ul", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, e *colly.HTMLElement) {
			// day
			day = append(day, e.ChildText(".day"))
			title = append(title, []string{})
			link = append(link, []string{})
			e.ForEach("a", func(_ int, e *colly.HTMLElement) {
				// title
				if e.Text != "" {
					title[i] = append(title[i], strings.Replace(e.Text, " ", "", -1))
				}
				// link
				if strings.Index(e.Attr("href"), "app") == -1 {
					link[i] = append(link[i],
						os.Getenv("APP_DOMAIN")+
							os.Getenv("APP_CALENDAR")+
							strings.Replace(e.Attr("href"), "./", "", 1),
					)
				}
			})
		})
	})

	c.Visit(os.Getenv("APP_DOMAIN") + os.Getenv("APP_CALENDAR") + os.Getenv("APP_CALENDAR_LIST") + "&cal_yy=" + year + "&cal_mm=" + month)

	calendar := []*domain.Calendar{}
	for i := 0; i < len(day); i++ {
		if day[i] != "" {
			calendar = append(calendar, &domain.Calendar{
				Day:   day[i],
				Plans: []domain.Plans{},
			})
			for j := 0; j < len(title[i]); j++ {
				calendar[i].Plans = append(calendar[i].Plans, domain.Plans{
					Title: title[i][j],
					Link:  link[i][j],
				})
			}
		}
	}

	// Plans が空の場合削除
	for i := 0; i < len(calendar); i++ {
		if len(calendar[i].Plans) == 0 {
			calendar = append(calendar[:i], calendar[i+1:]...)
			i--
		}
	}

	return calendar, nil
}
