package usecase

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type AttendanceUsecase interface {
	Get(year, month string, user *domain.User) ([]*domain.Attendance, error)
}

type attendanceUsecase struct {
	attendanceRepository repository.AttendanceRepository
}

func NewAttendanceUsecase(attendanceRepository repository.AttendanceRepository) AttendanceUsecase {
	return &attendanceUsecase{attendanceRepository: attendanceRepository}
}

func (u *attendanceUsecase) Get(year, month string, user *domain.User) ([]*domain.Attendance, error) {
	getAttendance, err := u.attendanceRepository.Get(user)
	if err != nil {
		return nil, err
	}

	return getAttendance, nil
}
