// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/yumekiti/eccSchoolApp-api/domain"
)

// AttendanceUsecase is an autogenerated mock type for the AttendanceUsecase type
type AttendanceUsecase struct {
	mock.Mock
}

// Get provides a mock function with given fields: user
func (_m *AttendanceUsecase) Get(user *domain.User) ([]*domain.Attendance, error) {
	ret := _m.Called(user)

	var r0 []*domain.Attendance
	if rf, ok := ret.Get(0).(func(*domain.User) []*domain.Attendance); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Attendance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAttendanceUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAttendanceUsecase creates a new instance of AttendanceUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAttendanceUsecase(t mockConstructorTestingTNewAttendanceUsecase) *AttendanceUsecase {
	mock := &AttendanceUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
