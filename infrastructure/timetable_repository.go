package infrastructure

import (
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type TimetableRepository struct{}

func NewTimetableRepository() repository.TimetableRepository {
	return &TimetableRepository{}
}

func (r *TimetableRepository) Get(week int, user *domain.User) (*domain.Timetable, error) {
	c, token := config.FalconLogin(user)

	// 返す値の初期化
	var date string
	var weekday string
	var period []string
	var subjectTitle []string
	var classroom []string
	var teacher []string

	// 取得時に必要な値を初期化
	var action string    // リクエストの種類を格納する変数
	var viewstate string // viewstateを格納する変数
	var links []string   // リンクを格納する配列
	weeks := []map[string]string{
		{"english": "Monday", "japanese": "月"},
		{"english": "Tuesday", "japanese": "火"},
		{"english": "Wednesday", "japanese": "水"},
		{"english": "Thursday", "japanese": "木"},
		{"english": "Friday", "japanese": "金"},
		{"english": "Saturday", "japanese": "土"},
		{"english": "Sunday", "japanese": "日"},
	}

	// viewstate を取得
	c.OnHTML("input", func(e *colly.HTMLElement) {
		if e.Attr("name") == "__VIEWSTATE" {
			viewstate = e.Attr("value")
		}
	})

	// action を取得
	c.OnHTML("form", func(e *colly.HTMLElement) {
		// action を取得
		action = e.Attr("action")
	})

	c.Visit(os.Getenv("FALCON_DOMAIN") + "/eccmo/(S(" + token + "))/MO0400/MO0400_01.aspx")

	c.OnHTML("form", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Text != "メインメニュー" && e.Text != "戻る" {
				subjectTitle = append(subjectTitle, e.Text)
				links = append(links, e.Attr("href"))
			}
		})
	})

	c.Post(os.Getenv("FALCON_DOMAIN")+"/eccmo/(S("+token+"))/MO0400/"+action,
		map[string]string{
			"__VIEWSTATE":     viewstate,
			"__EVENTTARGET":   "lnk" + weeks[(week - 1)]["english"],
			"__EVENTARGUMENT": "f" + weeks[(week - 1)]["english"],
			"txtDate":         "",
		})

	weekday = weeks[(week - 1)]["japanese"]

	for index, link := range links {
		if index == 0 {
			c.OnHTML("form", func(e *colly.HTMLElement) {
				date = strings.Split(strings.Split(e.Text, "日　付 : ")[1], "\n")[0]
				period = append(period, strings.Split(strings.Split(e.Text, "時　限 : ")[1], "\n")[0])
				classroom = append(classroom, strings.Split(strings.Split(e.Text, "教　室 : ")[1], "\n")[0])
				teacher = append(teacher, strings.Split(strings.Split(e.Text, "講　師 : ")[1], "\n")[0])
			})
		}

		c.Visit(os.Getenv("FALCON_DOMAIN") + strings.ReplaceAll(link, " ", ""))
	}

	timetable := domain.Timetable{}
	timetable.Date = date
	timetable.Weekday = weekday
	for index, _ := range subjectTitle {
		timetable.Timetable = append(timetable.Timetable, domain.TimetableDetail{
			Period:       period[index],
			SubjectTitle: subjectTitle[index],
			Classroom:    classroom[index],
			Teacher:      teacher[index],
		})
	}

	return &timetable, nil
}
