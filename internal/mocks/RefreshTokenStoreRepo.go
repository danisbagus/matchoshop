// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	mock "github.com/stretchr/testify/mock"
)

// RefreshTokenStoreRepo is an autogenerated mock type for the RefreshTokenStoreRepo type
type RefreshTokenStoreRepo struct {
	mock.Mock
}

// CheckRefreshToken provides a mock function with given fields: refreshToken
func (_m *RefreshTokenStoreRepo) CheckRefreshToken(refreshToken string) (bool, *errs.AppError) {
	ret := _m.Called(refreshToken)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(refreshToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(refreshToken)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Insert provides a mock function with given fields: refreshToken
func (_m *RefreshTokenStoreRepo) Insert(refreshToken string) *errs.AppError {
	ret := _m.Called(refreshToken)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(string) *errs.AppError); ok {
		r0 = rf(refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

type mockConstructorTestingTNewRefreshTokenStoreRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewRefreshTokenStoreRepo creates a new instance of RefreshTokenStoreRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRefreshTokenStoreRepo(t mockConstructorTestingTNewRefreshTokenStoreRepo) *RefreshTokenStoreRepo {
	mock := &RefreshTokenStoreRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
