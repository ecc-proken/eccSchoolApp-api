package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttendance(t *testing.T) {
	attendance := Attendance{
		Title:    "コンセプトワーク演習",
		Rate:     "100％",
		Absence:  "０",
		Lateness: "０",
	}

	assert.Equal(t, "コンセプトワーク演習", attendance.Title)
	assert.Equal(t, "100％", attendance.Rate)
	assert.Equal(t, "０", attendance.Absence)
	assert.Equal(t, "０", attendance.Lateness)
}