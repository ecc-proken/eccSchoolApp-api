// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/yumekiti/eccSchoolApp-api/domain"
)

// NewsOnlyUsecase is an autogenerated mock type for the NewsOnlyUsecase type
type NewsOnlyUsecase struct {
	mock.Mock
}

// Get provides a mock function with given fields: id, user
func (_m *NewsOnlyUsecase) Get(id string, user *domain.User) (*domain.NewsOnly, error) {
	ret := _m.Called(id, user)

	var r0 *domain.NewsOnly
	if rf, ok := ret.Get(0).(func(string, *domain.User) *domain.NewsOnly); ok {
		r0 = rf(id, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.NewsOnly)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *domain.User) error); ok {
		r1 = rf(id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewNewsOnlyUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewsOnlyUsecase creates a new instance of NewsOnlyUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewsOnlyUsecase(t mockConstructorTestingTNewNewsOnlyUsecase) *NewsOnlyUsecase {
	mock := &NewsOnlyUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
