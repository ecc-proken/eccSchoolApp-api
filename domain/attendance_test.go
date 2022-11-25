package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttendance(t *testing.T) {
	attendance := Attendance{
		Title:    "AIシステム開発演習Ⅰ",
		Rate:     "100％",
		Count:    "0",
		Absence:  "0",
		Lateness: "0",
	}

	assert.Equal(t, "AIシステム開発演習Ⅰ", attendance.Title)
	assert.Equal(t, "100％", attendance.Rate)
	assert.Equal(t, "0", attendance.Count)
	assert.Equal(t, "0", attendance.Absence)
	assert.Equal(t, "0", attendance.Lateness)
}
