package repository

import "github.com/yumekiti/eccSchoolApp-api/domain"

type TimetableRepository interface {
	Get(week int, user *domain.User) (*domain.Timetable, error)
}
