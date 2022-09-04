package repository

import "github.com/yumekiti/eccSchoolApp-api/domain"

type TimetableRepository interface {
	Get(week string, timetable *domain.Timetable) (*domain.Timetable, error)
}
