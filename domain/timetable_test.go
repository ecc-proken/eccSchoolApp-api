package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimetable(t *testing.T) {
	timetable := Timetable{
		Date:     "2022/11/21",
		Weekday:  "月",
		Timetable: []TimetableDetail{
			{
				Period: "1限",
				SubjectTitle: "セキュリティ演習_B",
				Classroom: "1405",
				Teacher: "教員の名前",
			},
		},
	}

	assert.Equal(t, "2022/11/21", timetable.Date)
	assert.Equal(t, "月", timetable.Weekday)
	assert.Equal(t, "1限", timetable.Timetable[0].Period)
	assert.Equal(t, "セキュリティ演習_B", timetable.Timetable[0].SubjectTitle)
	assert.Equal(t, "1405", timetable.Timetable[0].Classroom)
	assert.Equal(t, "教員の名前", timetable.Timetable[0].Teacher)
}
