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
			"pw":       user.Passwd,
		})
	if err != nil {
		log.Fatal(err)
	}

	return c
}
