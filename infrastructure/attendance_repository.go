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

	// 返す値の初期化
	var title []string
	var rate []string
	var count []string
	var absence []string
	var lateness []string

	// 取得時に必要な値を初期化
	var tmp [][]string   // 授業名と出席率を格納する配列
	var action string    // リクエストの種類を格納する変数
	var viewstate string // viewstateを格納する変数

	// viewstate を取得
	c.OnHTML("input", func(e *colly.HTMLElement) {
		if e.Attr("name") == "__VIEWSTATE" {
			viewstate = e.Attr("value")
		}
	})

	// action を取得
	c.OnHTML("form", func(e *colly.HTMLElement) {
		action = e.Attr("action")
	})

	// tmp を取得
	c.OnHTML("form", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Text != "メインメニュー" && e.Text != "欠席詳細" && e.Text != "戻る" {
				tmp = append(tmp, strings.Split(e.Text, " "))
			}
		})
	})

	c.Visit(os.Getenv("FALCON_DOMAIN") + "/eccmo/(S(" + token + "))/MO0500/MO0500_01.aspx?mode=1")

	// tmp 空白削除
	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp[i]); j++ {
			if tmp[i][j] == "" {
				tmp[i] = append(tmp[i][:j], tmp[i][j+1:]...)
			}
		}
	}

	for i, v := range tmp {
		// title
		title = append(title, v[0])
		// rate
		if len(v) == 2 {
			rate = append(rate, v[1])
		} else {
			rate = append(rate, "0%")
		}

		flag := false
		c.OnHTML("form", func(e *colly.HTMLElement) {
			if flag {
				return
			}

			// absence lateness
			if len(v) == 2 {
				count = append(count, strings.Split(strings.Split(e.Text, "出　席:")[1], "\n")[0])
				absence = append(absence, strings.Split(strings.Split(e.Text, "欠　席:")[1], "\n")[0])
				lateness = append(lateness, strings.Split(strings.Split(e.Text, "遅　刻:")[1], "\n")[0])
			} else {
				count = append(count, "0")
				absence = append(absence, "0")
				lateness = append(lateness, "0")
			}

			flag = true
		})

		err := c.Post(os.Getenv("FALCON_DOMAIN")+"/eccmo/(S("+token+"))/MO0500/"+action,
			map[string]string{
				"__VIEWSTATE":     viewstate,
				"__EVENTTARGET":   "lstSyussekiRitsu",
				"__EVENTARGUMENT": fmt.Sprint(i),
			})
		if err != nil {
			log.Fatal(err)
		}
	}

	// 返す値から attendance を作成
	attendance := []*domain.Attendance{}
	for i, v := range title {
		attendance = append(attendance, &domain.Attendance{
			Title:    v,
			Rate:     rate[i],
			Count:    count[i],
			Absence:  absence[i],
			Lateness: lateness[i],
		})
	}

	return attendance, nil
}
