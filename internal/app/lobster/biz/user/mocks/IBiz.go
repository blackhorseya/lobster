// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/lobster/internal/pkg/contextx"
	mock "github.com/stretchr/testify/mock"

	pb "github.com/blackhorseya/lobster/internal/pkg/pb"
)

// IBiz is an autogenerated mock type for the IBiz type
type IBiz struct {
	mock.Mock
}

// GetInfoByAccessToken provides a mock function with given fields: ctx, token
func (_m *IBiz) GetInfoByAccessToken(ctx contextx.Contextx, token string) (*pb.Profile, error) {
	ret := _m.Called(ctx, token)

	var r0 *pb.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *pb.Profile); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInfoByEmail provides a mock function with given fields: ctx, email
func (_m *IBiz) GetInfoByEmail(ctx contextx.Contextx, email string) (*pb.Profile, error) {
	ret := _m.Called(ctx, email)

	var r0 *pb.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *pb.Profile); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInfoByID provides a mock function with given fields: ctx, id
func (_m *IBiz) GetInfoByID(ctx contextx.Contextx, id string) (*pb.Profile, error) {
	ret := _m.Called(ctx, id)

	var r0 *pb.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *pb.Profile); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Profile)
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

// Login provides a mock function with given fields: ctx, email, token
func (_m *IBiz) Login(ctx contextx.Contextx, email string, token string) (*pb.Profile, error) {
	ret := _m.Called(ctx, email, token)

	var r0 *pb.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, string) *pb.Profile); ok {
		r0 = rf(ctx, email, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, string) error); ok {
		r1 = rf(ctx, email, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx, _a1
func (_m *IBiz) Logout(ctx contextx.Contextx, _a1 *pb.Profile) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *pb.Profile) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Signup provides a mock function with given fields: ctx, email, token
func (_m *IBiz) Signup(ctx contextx.Contextx, email string, token string) (*pb.Profile, error) {
	ret := _m.Called(ctx, email, token)

	var r0 *pb.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, string) *pb.Profile); ok {
		r0 = rf(ctx, email, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, string) error); ok {
		r1 = rf(ctx, email, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}