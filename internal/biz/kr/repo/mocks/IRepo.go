// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/lobster/internal/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	okr "github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *IRepo) Delete(ctx contextx.Contextx, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryByID provides a mock function with given fields: ctx, id
func (_m *IRepo) QueryByID(ctx contextx.Contextx, id string) (*okr.KeyResult, error) {
	ret := _m.Called(ctx, id)

	var r0 *okr.KeyResult
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *okr.KeyResult); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*okr.KeyResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryList provides a mock function with given fields: ctx, offset, limit
func (_m *IRepo) QueryList(ctx contextx.Contextx, offset int, limit int) ([]*okr.KeyResult, error) {
	ret := _m.Called(ctx, offset, limit)

	var r0 []*okr.KeyResult
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int, int) []*okr.KeyResult); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*okr.KeyResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int, int) error); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
