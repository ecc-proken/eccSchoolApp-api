package usecase

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type TimetableUsecase interface {
	Get(week string, timetable *domain.Timetable) (*domain.Timetable, error)
}

type timetableUsecase struct {
	timetableRepository repository.TimetableRepository
}

func NewTimetableUsecase(timetableRepository repository.TimetableRepository) TimetableUsecase {
	return &timetableUsecase{timetableRepository: timetableRepository}
}

func (u *timetableUsecase) Get(week string, timetable *domain.Timetable) (*domain.Timetable, error) {
	getTimetable, err := u.timetableRepository.Get(week, timetable)
	if err != nil {
		return nil, err
	}

	return getTimetable, nil
}
