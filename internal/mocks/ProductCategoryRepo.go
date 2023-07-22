// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	domain "github.com/danisbagus/matchoshop/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// ProductCategoryRepo is an autogenerated mock type for the ProductCategoryRepo type
type ProductCategoryRepo struct {
	mock.Mock
}

// CheckByID provides a mock function with given fields: productCategoryID
func (_m *ProductCategoryRepo) CheckByID(productCategoryID int64) (bool, *errs.AppError) {
	ret := _m.Called(productCategoryID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64) bool); ok {
		r0 = rf(productCategoryID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(productCategoryID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// CheckByIDAndName provides a mock function with given fields: productCategoryID, name
func (_m *ProductCategoryRepo) CheckByIDAndName(productCategoryID int64, name string) (bool, *errs.AppError) {
	ret := _m.Called(productCategoryID, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64, string) bool); ok {
		r0 = rf(productCategoryID, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64, string) *errs.AppError); ok {
		r1 = rf(productCategoryID, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// CheckByName provides a mock function with given fields: name
func (_m *ProductCategoryRepo) CheckByName(name string) (bool, *errs.AppError) {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Delete provides a mock function with given fields: productCategoryID
func (_m *ProductCategoryRepo) Delete(productCategoryID int64) *errs.AppError {
	ret := _m.Called(productCategoryID)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int64) *errs.AppError); ok {
		r0 = rf(productCategoryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ProductCategoryRepo) GetAll() ([]domain.ProductCategory, *errs.AppError) {
	ret := _m.Called()

	var r0 []domain.ProductCategory
	if rf, ok := ret.Get(0).(func() []domain.ProductCategory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ProductCategory)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func() *errs.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// GetAllByProductID provides a mock function with given fields: productID
func (_m *ProductCategoryRepo) GetAllByProductID(productID int64) ([]domain.ProductCategory, *errs.AppError) {
	ret := _m.Called(productID)

	var r0 []domain.ProductCategory
	if rf, ok := ret.Get(0).(func(int64) []domain.ProductCategory); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ProductCategory)
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

// GetOneByID provides a mock function with given fields: productCategoryID
func (_m *ProductCategoryRepo) GetOneByID(productCategoryID int64) (*domain.ProductCategory, *errs.AppError) {
	ret := _m.Called(productCategoryID)

	var r0 *domain.ProductCategory
	if rf, ok := ret.Get(0).(func(int64) *domain.ProductCategory); ok {
		r0 = rf(productCategoryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ProductCategory)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(productCategoryID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data
func (_m *ProductCategoryRepo) Insert(data *domain.ProductCategory) (*domain.ProductCategory, *errs.AppError) {
	ret := _m.Called(data)

	var r0 *domain.ProductCategory
	if rf, ok := ret.Get(0).(func(*domain.ProductCategory) *domain.ProductCategory); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ProductCategory)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(*domain.ProductCategory) *errs.AppError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: productCategoryID, data
func (_m *ProductCategoryRepo) Update(productCategoryID int64, data *domain.ProductCategory) *errs.AppError {
	ret := _m.Called(productCategoryID, data)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int64, *domain.ProductCategory) *errs.AppError); ok {
		r0 = rf(productCategoryID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

type mockConstructorTestingTNewProductCategoryRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductCategoryRepo creates a new instance of ProductCategoryRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductCategoryRepo(t mockConstructorTestingTNewProductCategoryRepo) *ProductCategoryRepo {
	mock := &ProductCategoryRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
