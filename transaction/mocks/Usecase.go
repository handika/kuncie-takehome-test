// Code generated by mockery v1.0.0
package mocks

import (
	context "context"

	models "github.com/handika/kuncie-takehome-test/models"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) GetByID(ctx context.Context, id int64) (*models.Transaction, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) Store(_a0 context.Context, _a1 *models.Transaction) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Transaction) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, ar
func (_m *Usecase) Update(ctx context.Context, ar *models.Transaction) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Transaction) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
