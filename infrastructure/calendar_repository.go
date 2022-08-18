package infrastructure

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type CalendarRepository struct{}

func NewCalendarRepository() repository.CalendarRepository {
	return &CalendarRepository{}
}

func (r *CalendarRepository) Get(year, month string, user *domain.User) ([]*domain.Calendar, error) {
	return []*domain.Calendar{}, nil
}
