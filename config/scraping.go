package config

import "github.com/gocolly/colly"

func NewColly() *colly.Collector {
	return colly.NewCollector()
}
