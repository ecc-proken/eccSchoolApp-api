// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/yumekiti/eccSchoolApp-api/domain"
)

// NewsOnlyRepository is an autogenerated mock type for the NewsOnlyRepository type
type NewsOnlyRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: id, user
func (_m *NewsOnlyRepository) Get(id string, user *domain.User) (*domain.NewsOnly, error) {
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

type mockConstructorTestingTNewNewsOnlyRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewsOnlyRepository creates a new instance of NewsOnlyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewsOnlyRepository(t mockConstructorTestingTNewNewsOnlyRepository) *NewsOnlyRepository {
	mock := &NewsOnlyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
