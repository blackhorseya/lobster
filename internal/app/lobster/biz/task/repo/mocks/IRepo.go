// Code generated by mockery (devel). DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	mock "github.com/stretchr/testify/mock"

	todo "github.com/blackhorseya/lobster/internal/pkg/entity/todo"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, userID
func (_m *IRepo) Count(ctx contextx.Contextx, userID int64) (int, error) {
	ret := _m.Called(ctx, userID)

	var r0 int
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64) int); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64) error); ok {
		r1 = rf(ctx, userID)
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
func (_m *IRepo) Delete(ctx contextx.Contextx, id int64) (int, error) {
	ret := _m.Called(ctx, id)

	var r0 int
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64) int); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, userID, offset, limit
func (_m *IRepo) List(ctx contextx.Contextx, userID int64, offset int, limit int) ([]*todo.Task, error) {
	ret := _m.Called(ctx, userID, offset, limit)

	var r0 []*todo.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64, int, int) []*todo.Task); ok {
		r0 = rf(ctx, userID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*todo.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64, int, int) error); ok {
		r1 = rf(ctx, userID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryByID provides a mock function with given fields: ctx, userID, id
func (_m *IRepo) QueryByID(ctx contextx.Contextx, userID int64, id int64) (*todo.Task, error) {
	ret := _m.Called(ctx, userID, id)

	var r0 *todo.Task
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64, int64) *todo.Task); ok {
		r0 = rf(ctx, userID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*todo.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64, int64) error); ok {
		r1 = rf(ctx, userID, id)
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
