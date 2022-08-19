package config

import (
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/yumekiti/eccSchoolApp-api/domain"
)

// ログイン処理
func ECCLogin(user *domain.User) *colly.Collector {
	c := colly.NewCollector()

	err := c.Post(os.Getenv("APP_DOMAIN")+os.Getenv("APP_LOGIN"),
		map[string]string{
			"c":        "login_2",
			"flg_auto": "1",
			"token_a":  "",
			"id":       user.Id,
			"pw":       user.Password,
		})
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// func FalconLogin(user *domain.User) (*colly.Collector, string) {
// 	// リダイレクト先を取得
// 	target_url := "https://falcon.ecc.ac.jp/eccmo/mo0100/mo0100_01.aspx"
// 	req, _ := http.NewRequest("HEAD", target_url, nil)
// 	resp, _ := http.DefaultTransport.RoundTrip(req)
// 	defer resp.Body.Close()
// 	location := resp.Header["Location"][0]

// 	// token を取得
// 	token := strings.Split(strings.Split(location, "(")[2], ")")[0]

// 	c := colly.NewCollector()

// 	// 初期化
// 	action := ""
// 	viewstate := ""

// 	// formのactionを取得
// 	c.OnHTML("form", func(e *colly.HTMLElement) {
// 		action = e.Attr("action")
// 	})

// 	// inputのvalueを取得
// 	c.OnHTML("input", func(e *colly.HTMLElement) {
// 		if e.Attr("name") == "__VIEWSTATE" {
// 			viewstate = e.Attr("value")
// 		}
// 	})

// 	c.Visit("https://falcon.ecc.ac.jp/eccmo/(S(" + token + "))/MO0100/MO0100_01.aspx")

// 	err := c.Post("https://falcon.ecc.ac.jp/eccmo/(S("+token+"))/MO0100/"+action,
// 		map[string]string{
// 			"__VIEWSTATE":     viewstate,
// 			"__EVENTTARGET":   "",
// 			"__EVENTARGUMENT": "",
// 			"txtUserId":       user.Id,
// 			"txtPassword":     user.Password,
// 		})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	c.OnResponse(func(r *colly.Response) {
// 		r.Request.Headers.Set("Referer", "https://falcon.ecc.ac.jp/eccmo/(S("+token+"))/MO0100/MO0100_01.aspx")
// 	})

// 	return c, token
// }
