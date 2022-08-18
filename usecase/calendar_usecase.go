package usecase

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type CalendarUsecase interface {
	Get(year, month string, user *domain.User) ([]*domain.Calendar, error)
}

type calendarUsecase struct {
	calendarRepository repository.CalendarRepository
}

func NewCalendarUsecase(calendarRepository repository.CalendarRepository) CalendarUsecase {
	return &calendarUsecase{calendarRepository: calendarRepository}
}

func (u *calendarUsecase) Get(year, month string, user *domain.User) ([]*domain.Calendar, error) {
	getCalendar, err := u.calendarRepository.Get(year, month, user)
	if err != nil {
		return nil, err
	}

	return getCalendar, nil
}
