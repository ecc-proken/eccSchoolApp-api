package infrastructure

import (
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type AttendanceRepository struct{}

func NewAttendanceRepository() repository.AttendanceRepository {
	return &AttendanceRepository{}
}

func (r *AttendanceRepository) Get(user *domain.User) ([]*domain.Attendance, error) {
	return []*domain.Attendance{}, nil
}
