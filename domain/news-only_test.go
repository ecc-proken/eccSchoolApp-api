package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewsOnly(t *testing.T) {
	newsOnly := NewsOnly{
		Title:      "【留学生と日本文化にチャレンジ　かるたで遊ぼう！】",
		Body:       "ECCでは日本人学生と留学生の交流会　BUDDY　PROGRAMを開催しています。<br/><a href=\"https://example.com\" target=\"_blank\">https://example.com</a><br/><br/>皆さんのご参加をお待ちしております！",
		Date:       "2022/11/18",
		Tag:        "学校からの連絡",
		Attachment: []string{"https://example.com"},
	}

	assert.Equal(t, "【留学生と日本文化にチャレンジ　かるたで遊ぼう！】", newsOnly.Title)
	assert.Equal(t, "ECCでは日本人学生と留学生の交流会　BUDDY　PROGRAMを開催しています。<br/><a href=\"https://example.com\" target=\"_blank\">https://example.com</a><br/><br/>皆さんのご参加をお待ちしております！", newsOnly.Body)
	assert.Equal(t, "2022/11/18", newsOnly.Date)
	assert.Equal(t, "学校からの連絡", newsOnly.Tag)
	assert.Equal(t, "https://example.com", newsOnly.Attachment[0])
}
