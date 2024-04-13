// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	mock "github.com/stretchr/testify/mock"
)

// IRefreshTokenStoreRepository is an autogenerated mock type for the IRefreshTokenStoreRepository type
type IRefreshTokenStoreRepository struct {
	mock.Mock
}

// CheckRefreshToken provides a mock function with given fields: refreshToken
func (_m *IRefreshTokenStoreRepository) CheckRefreshToken(refreshToken string) (bool, *errs.AppError) {
	ret := _m.Called(refreshToken)

	var r0 bool
	var r1 *errs.AppError
	if rf, ok := ret.Get(0).(func(string) (bool, *errs.AppError)); ok {
		return rf(refreshToken)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(refreshToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

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
func (_m *IRefreshTokenStoreRepository) Insert(refreshToken string) *errs.AppError {
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

// NewIRefreshTokenStoreRepository creates a new instance of IRefreshTokenStoreRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRefreshTokenStoreRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRefreshTokenStoreRepository {
	mock := &IRefreshTokenStoreRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
