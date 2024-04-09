// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	domain "github.com/danisbagus/matchoshop/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// IProductProductCategoryRepository is an autogenerated mock type for the IProductProductCategoryRepository type
type IProductProductCategoryRepository struct {
	mock.Mock
}

// BulkInsert provides a mock function with given fields: data
func (_m *IProductProductCategoryRepository) BulkInsert(data []domain.ProductProductCategory) *errs.AppError {
	ret := _m.Called(data)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func([]domain.ProductProductCategory) *errs.AppError); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// DeleteAll provides a mock function with given fields: productID
func (_m *IProductProductCategoryRepository) DeleteAll(productID int64) *errs.AppError {
	ret := _m.Called(productID)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int64) *errs.AppError); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

type mockConstructorTestingTNewIProductProductCategoryRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIProductProductCategoryRepository creates a new instance of IProductProductCategoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIProductProductCategoryRepository(t mockConstructorTestingTNewIProductProductCategoryRepository) *IProductProductCategoryRepository {
	mock := &IProductProductCategoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
