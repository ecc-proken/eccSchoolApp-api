package config

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/yumekiti/eccSchoolApp-api/domain"
)

func ECCLogin(user *domain.User) *colly.Collector {
	c := colly.NewCollector()

	// ログイン処理
	err := c.Post(os.Getenv("APP_DOMAIN")+os.Getenv("APP_LOGIN"),
		map[string]string{
			"c":        "login_2",
			"flg_auto": "1",
			"token_a":  "",
			"id":       user.ID,
			"pw":       user.Password,
		})
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func FalconLogin(user *domain.User) (*colly.Collector, string) {
	// リダイレクト先のtokenを取得
	target_url := os.Getenv("FALCON") + "/eccmo/mo0100/mo0100_01.aspx"
	req, _ := http.NewRequest("HEAD", target_url, nil)
	resp, _ := http.DefaultTransport.RoundTrip(req)
	defer resp.Body.Close()
	location := resp.Header["Location"][0]

	// tokenを取得
	token := strings.Split(strings.Split(location, "(")[2], ")")[0]

	// ログイン時に必要なデータ取得
	c := colly.NewCollector()

	var action string
	var viewstate string

	// action を取得
	c.OnHTML("form", func(e *colly.HTMLElement) {
		action = e.Attr("action")
	})

	// viewstate を取得
	c.OnHTML("input", func(e *colly.HTMLElement) {
		if e.Attr("name") == "__VIEWSTATE" {
			viewstate = e.Attr("value")
		}
	})

	c.Visit(os.Getenv("FALCON") + "/eccmo/(S(" + token + "))/MO0100/MO0100_01.aspx")

	// ログイン処理
	err := c.Post(os.Getenv("FALCON")+"/eccmo/(S("+token+"))/MO0100/"+action,
		map[string]string{
			"__VIEWSTATE":     viewstate,
			"__EVENTTARGET":   "",
			"__EVENTARGUMENT": "",
			"txtUserId":       user.ID,
			"txtPassword":     user.Password,
			"btnLogin":        "",
		})
	if err != nil {
		log.Fatal(err)
	}

	return c, token
}
