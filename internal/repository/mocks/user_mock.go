// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	domain "github.com/danisbagus/matchoshop/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// CreateUserCustomer provides a mock function with given fields: data
func (_m *IUserRepository) CreateUserCustomer(data *domain.UserModel) (*domain.UserModel, *errs.AppError) {
	ret := _m.Called(data)

	var r0 *domain.UserModel
	var r1 *errs.AppError
	if rf, ok := ret.Get(0).(func(*domain.UserModel) (*domain.UserModel, *errs.AppError)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*domain.UserModel) *domain.UserModel); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.UserModel) *errs.AppError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID
func (_m *IUserRepository) Delete(userID int64) *errs.AppError {
	ret := _m.Called(userID)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int64) *errs.AppError); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// FindOne provides a mock function with given fields: email
func (_m *IUserRepository) FindOne(email string) (*domain.UserModel, *errs.AppError) {
	ret := _m.Called(email)

	var r0 *domain.UserModel
	var r1 *errs.AppError
	if rf, ok := ret.Get(0).(func(string) (*domain.UserModel, *errs.AppError)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.UserModel); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// FindOneById provides a mock function with given fields: userID
func (_m *IUserRepository) FindOneById(userID int64) (*domain.UserModel, *errs.AppError) {
	ret := _m.Called(userID)

	var r0 *domain.UserModel
	var r1 *errs.AppError
	if rf, ok := ret.Get(0).(func(int64) (*domain.UserModel, *errs.AppError)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int64) *domain.UserModel); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *IUserRepository) GetAll() ([]domain.UserDetail, *errs.AppError) {
	ret := _m.Called()

	var r0 []domain.UserDetail
	var r1 *errs.AppError
	if rf, ok := ret.Get(0).(func() ([]domain.UserDetail, *errs.AppError)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.UserDetail); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UserDetail)
		}
	}

	if rf, ok := ret.Get(1).(func() *errs.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: userID, data
func (_m *IUserRepository) Update(userID int64, data *domain.UserModel) *errs.AppError {
	ret := _m.Called(userID, data)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int64, *domain.UserModel) *errs.AppError); ok {
		r0 = rf(userID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
