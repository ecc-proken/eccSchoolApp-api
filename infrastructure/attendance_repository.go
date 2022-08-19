package infrastructure

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type AttendanceRepository struct{}

func NewAttendanceRepository() repository.AttendanceRepository {
	return &AttendanceRepository{}
}

func (r *AttendanceRepository) Get(user *domain.User) ([]*domain.Attendance, error) {
	c, token := config.FalconLogin(user)

	// 初期化
	title := []string{}
	rate := []string{}
	absence := []string{}
	lateness := []string{}

	tmp := [][]string{}
	action := ""
	viewstate := ""
	venttarget := "lstSyussekiRitsu"

	c.OnHTML("input", func(e *colly.HTMLElement) {
		if e.Attr("name") == "__VIEWSTATE" {
			viewstate = e.Attr("value")
		}
	})

	c.OnHTML("form", func(e *colly.HTMLElement) {
		action = e.Attr("action")
	})

	c.OnHTML("form", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Text == "戻る" {
				return
			}
			if len(strings.Split(e.Text, " ")) == 2 {
				tmp = append(tmp, strings.Split(e.Text, " "))
			}
		})
	})

	c.Visit(os.Getenv("FALCON") + "/eccmo/(S(" + token + "))/MO0500/MO0500_01.aspx?mode=1")

	for i, v := range tmp {
		// title
		title = append(title, v[0])
		// rate
		rate = append(rate, v[1])

		flag := false
		c.OnHTML("form", func(e *colly.HTMLElement) {
			if flag {
				return
			}

			absence = append(absence, strings.Split(e.Text, "欠　席:")[1][:1])
			lateness = append(lateness, strings.Split(e.Text, "遅　刻:")[1][:1])

			flag = true
		})

		err := c.Post(os.Getenv("FALCON")+"/eccmo/(S("+token+"))/MO0500/"+action,
			map[string]string{
				"__VIEWSTATE":     viewstate,
				"__EVENTTARGET":   venttarget,
				"__EVENTARGUMENT": fmt.Sprint(i),
			})
		if err != nil {
			log.Fatal(err)
		}
	}

	attendance := []*domain.Attendance{}
	for i, v := range title {
		attendance = append(attendance, &domain.Attendance{
			Title:    v,
			Rate:     rate[i],
			Absence:  absence[i],
			Lateness: lateness[i],
		})
	}

	return attendance, nil
}
