// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/yumekiti/eccSchoolApp-api/domain"
)

// SigninRepository is an autogenerated mock type for the SigninRepository type
type SigninRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *SigninRepository) Get(_a0 *domain.User) (*domain.Signin, error) {
	ret := _m.Called(_a0)

	var r0 *domain.Signin
	if rf, ok := ret.Get(0).(func(*domain.User) *domain.Signin); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Signin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSigninRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewSigninRepository creates a new instance of SigninRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSigninRepository(t mockConstructorTestingTNewSigninRepository) *SigninRepository {
	mock := &SigninRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}