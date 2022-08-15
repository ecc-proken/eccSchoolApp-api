package entity

type Timetable struct {
	Date      string
	Weekday   string
	Timetable []TimetableDetail
}

type TimetableDetail struct {
	Period       string
	SubjectTitle string
	Classroom    string
	Teacher      string
}
