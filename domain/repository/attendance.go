package repository

import "github.com/yumekiti/eccSchoolApp-api/domain"

type AttendanceRepository interface {
	Get(user *domain.User) ([]*domain.Attendance, error)
}
