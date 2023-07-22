// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	domain "github.com/danisbagus/matchoshop/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// ReviewRepo is an autogenerated mock type for the ReviewRepo type
type ReviewRepo struct {
	mock.Mock
}

// GetAllByProductID provides a mock function with given fields: productID
func (_m *ReviewRepo) GetAllByProductID(productID int64) ([]domain.Review, *errs.AppError) {
	ret := _m.Called(productID)

	var r0 []domain.Review
	if rf, ok := ret.Get(0).(func(int64) []domain.Review); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Review)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(productID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// GetOneByUserIDAndProductID provides a mock function with given fields: userID, productID
func (_m *ReviewRepo) GetOneByUserIDAndProductID(userID int64, productID int64) (*domain.Review, *errs.AppError) {
	ret := _m.Called(userID, productID)

	var r0 *domain.Review
	if rf, ok := ret.Get(0).(func(int64, int64) *domain.Review); ok {
		r0 = rf(userID, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Review)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64, int64) *errs.AppError); ok {
		r1 = rf(userID, productID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Insert provides a mock function with given fields: form
func (_m *ReviewRepo) Insert(form *domain.Review) *errs.AppError {
	ret := _m.Called(form)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(*domain.Review) *errs.AppError); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// Update provides a mock function with given fields: form
func (_m *ReviewRepo) Update(form *domain.Review) *errs.AppError {
	ret := _m.Called(form)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(*domain.Review) *errs.AppError); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

type mockConstructorTestingTNewReviewRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewReviewRepo creates a new instance of ReviewRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReviewRepo(t mockConstructorTestingTNewReviewRepo) *ReviewRepo {
	mock := &ReviewRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}