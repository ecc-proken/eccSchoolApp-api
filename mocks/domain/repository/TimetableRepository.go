// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/yumekiti/eccSchoolApp-api/domain"
)

// TimetableRepository is an autogenerated mock type for the TimetableRepository type
type TimetableRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: week, user
func (_m *TimetableRepository) Get(week int, user *domain.User) (*domain.Timetable, error) {
	ret := _m.Called(week, user)

	var r0 *domain.Timetable
	if rf, ok := ret.Get(0).(func(int, *domain.User) *domain.Timetable); ok {
		r0 = rf(week, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Timetable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *domain.User) error); ok {
		r1 = rf(week, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTimetableRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTimetableRepository creates a new instance of TimetableRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTimetableRepository(t mockConstructorTestingTNewTimetableRepository) *TimetableRepository {
	mock := &TimetableRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}