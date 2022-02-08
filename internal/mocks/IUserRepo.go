// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	domain "github.com/danisbagus/matchoshop/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// IUserRepo is an autogenerated mock type for the IUserRepo type
type IUserRepo struct {
	mock.Mock
}

// FindOne provides a mock function with given fields: email
func (_m *IUserRepo) FindOne(email string) (*domain.User, *errs.AppError) {
	ret := _m.Called(email)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// GenerateAccessTokenAndRefreshToken provides a mock function with given fields: data
func (_m *IUserRepo) GenerateAccessTokenAndRefreshToken(data *domain.User) (string, string, *errs.AppError) {
	ret := _m.Called(data)

	var r0 string
	if rf, ok := ret.Get(0).(func(*domain.User) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*domain.User) string); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 *errs.AppError
	if rf, ok := ret.Get(2).(func(*domain.User) *errs.AppError); ok {
		r2 = rf(data)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*errs.AppError)
		}
	}

	return r0, r1, r2
}

// Verify provides a mock function with given fields: token
func (_m *IUserRepo) Verify(token string) *errs.AppError {
	ret := _m.Called(token)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(string) *errs.AppError); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}