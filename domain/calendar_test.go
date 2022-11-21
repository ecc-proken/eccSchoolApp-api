package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalendar(t *testing.T) {
	calendar := Calendar{
		Day: "9",
		Plans: []Plans{
			{
				Title: "情報処理技術者試験(IPA)",
				Link:  "https://example.com",
			},
		},
	}

	assert.Equal(t, "9", calendar.Day)
	assert.Equal(t, "情報処理技術者試験(IPA)", calendar.Plans[0].Title)
	assert.Equal(t, "https://example.com", calendar.Plans[0].Link)
}
