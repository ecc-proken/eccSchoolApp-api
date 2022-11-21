package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNews(t *testing.T) {
	news := News{
		ID:    "16687374031867150007",
		Title: "【留学生と日本文化にチャレンジ　かるたで遊ぼう！】",
		Date:  "2022/11/18",
		Tag:   "学校からの連絡",
		Link:  "https://example.com",
	}

	assert.Equal(t, "16687374031867150007", news.ID)
	assert.Equal(t, "【留学生と日本文化にチャレンジ　かるたで遊ぼう！】", news.Title)
	assert.Equal(t, "2022/11/18", news.Date)
	assert.Equal(t, "学校からの連絡", news.Tag)
	assert.Equal(t, "https://example.com", news.Link)
}