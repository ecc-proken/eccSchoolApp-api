package infrastructure

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type TimetableRepository struct{}

func NewTimetableRepository() repository.TimetableRepository {
	return &TimetableRepository{}
}

func (r *TimetableRepository) Get(week string, timetable *domain.Timetable) (*domain.Timetable, error) {
	return &domain.Timetable{
		Date:    "2021-01-01",
		Weekday: "月",
		Timetable: []domain.TimetableDetail{
			{
				Period:       "1",
				SubjectTitle: "国語",
				Classroom:    "1",
				Teacher:      "田中",
			},
			{
				Period:       "2",
				SubjectTitle: "数学",
				Classroom:    "2",
				Teacher:      "鈴木",
			},
		},
	}, nil
}
