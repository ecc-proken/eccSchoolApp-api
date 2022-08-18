package repository

import "github.com/yumekiti/eccSchoolApp-api/domain"

type CalendarRepository interface {
	Get(year, month string, user *domain.User) ([]*domain.Calendar, error)
}
