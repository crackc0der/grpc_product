// Code generated by mockery v2.43.1. DO NOT EDIT.

package category

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRepositoryInterface is an autogenerated mock type for the RepositoryInterface type
type MockRepositoryInterface struct {
	mock.Mock
}

// DeleteCategory provides a mock function with given fields: _a0, _a1
func (_m *MockRepositoryInterface) DeleteCategory(_a0 context.Context, _a1 int64) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCategory")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertCategory provides a mock function with given fields: _a0, _a1
func (_m *MockRepositoryInterface) InsertCategory(_a0 context.Context, _a1 *Category) (*Category, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for InsertCategory")
	}

	var r0 *Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Category) (*Category, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Category) *Category); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Category) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectCategories provides a mock function with given fields: _a0
func (_m *MockRepositoryInterface) SelectCategories(_a0 context.Context) ([]Category, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SelectCategories")
	}

	var r0 []Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]Category, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []Category); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCategory provides a mock function with given fields: _a0, _a1
func (_m *MockRepositoryInterface) UpdateCategory(_a0 context.Context, _a1 *Category) (*Category, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCategory")
	}

	var r0 *Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Category) (*Category, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Category) *Category); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Category) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockRepositoryInterface creates a new instance of MockRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
