package domain

// 時間割
type Timetable struct {
	Date      string            // 日付
	Weekday   string            // 曜日
	Timetable []TimetableDetail // 時間割の詳細
}

type TimetableDetail struct {
	Period       string
	SubjectTitle string
	Classroom    string
	Teacher      string
}
