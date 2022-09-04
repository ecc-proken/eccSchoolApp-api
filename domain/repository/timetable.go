package repository

import "github.com/yumekiti/eccSchoolApp-api/domain"

type TimetableRepository interface {
	Get(week string, user *domain.User) (*domain.Timetable, error)
}
