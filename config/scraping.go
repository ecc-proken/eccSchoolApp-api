package config

import (
	"log"
	"os"

	"github.com/gocolly/colly"
)

func NewColly() *colly.Collector {
	return colly.NewCollector()
}

// ログイン処理
func AppLogin(c *colly.Collector, id, pw string) *colly.Collector {
	err := c.Post(os.Getenv("APP_DOMAIN")+os.Getenv("APP_LOGIN"),
		map[string]string{
			"c":        "login_2",
			"flg_auto": "1",
			"token_a":  "",
			"id":       id,
			"pw":       pw,
		})
	if err != nil {
		log.Fatal(err)
	}

	return c
}
