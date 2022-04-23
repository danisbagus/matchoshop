// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	domain "github.com/danisbagus/matchoshop/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrderProductRepo is an autogenerated mock type for the OrderProductRepo type
type OrderProductRepo struct {
	mock.Mock
}

// GetAllByOrderID provides a mock function with given fields: orderUD
func (_m *OrderProductRepo) GetAllByOrderID(orderUD int64) ([]domain.OrderProduct, *errs.AppError) {
	ret := _m.Called(orderUD)

	var r0 []domain.OrderProduct
	if rf, ok := ret.Get(0).(func(int64) []domain.OrderProduct); ok {
		r0 = rf(orderUD)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.OrderProduct)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(orderUD)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}
