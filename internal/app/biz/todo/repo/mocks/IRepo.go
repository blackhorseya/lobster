// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/lobster/internal/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	todo "github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *IRepo) Count(ctx contextx.Contextx) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(contextx.Contextx) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, task
func (_m *IRepo) Create(ctx contextx.Contextx, task *todo.Task) (*todo.Task, error) {
	ret := _m.Called(ctx, task)

	var r0 *todo.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *todo.Task) *todo.Task); ok {
		r0 = rf(ctx, task)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *todo.Task) error); ok {
		r1 = rf(ctx, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *IRepo) Delete(ctx contextx.Contextx, id string) (int, error) {
	ret := _m.Called(ctx, id)

	var r0 int
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) int); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, offset, limit
func (_m *IRepo) List(ctx contextx.Contextx, offset int, limit int) ([]*todo.Task, error) {
	ret := _m.Called(ctx, offset, limit)

	var r0 []*todo.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int, int) []*todo.Task); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*todo.Task)
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

// Update provides a mock function with given fields: ctx, updated
func (_m *IRepo) Update(ctx contextx.Contextx, updated *todo.Task) (*todo.Task, error) {
	ret := _m.Called(ctx, updated)

	var r0 *todo.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *todo.Task) *todo.Task); ok {
		r0 = rf(ctx, updated)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *todo.Task) error); ok {
		r1 = rf(ctx, updated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
