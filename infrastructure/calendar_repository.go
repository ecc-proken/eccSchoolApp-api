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
	return []*domain.Calendar{
		{
			Day: "1",
			Plans: domain.Plans{
				Title: []string{"1日目の予定"},
				Link:  []string{"https://www.google.com/"},
			},
		},
		{
			Day: "2",
			Plans: domain.Plans{
				Title: []string{"2日目の予定"},
				Link:  []string{"https://www.google.com/"},
			},
		},
	}, nil
}
