// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mock

import (
	model "github.com/JesusG2000/hexsatisfaction/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Author is an autogenerated mock type for the Author type
type Author struct {
	mock.Mock
}

// Create provides a mock function with given fields: request
func (_m *Author) Create(request model.CreateAuthorRequest) (int, error) {
	ret := _m.Called(request)

	var r0 int
	if rf, ok := ret.Get(0).(func(model.CreateAuthorRequest) int); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.CreateAuthorRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: request
func (_m *Author) Delete(request model.DeleteAuthorRequest) (int, error) {
	ret := _m.Called(request)

	var r0 int
	if rf, ok := ret.Get(0).(func(model.DeleteAuthorRequest) int); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.DeleteAuthorRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *Author) FindAll() ([]model.Author, error) {
	ret := _m.Called()

	var r0 []model.Author
	if rf, ok := ret.Get(0).(func() []model.Author); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: request
func (_m *Author) FindByID(request model.IDAuthorRequest) (*model.Author, error) {
	ret := _m.Called(request)

	var r0 *model.Author
	if rf, ok := ret.Get(0).(func(model.IDAuthorRequest) *model.Author); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.IDAuthorRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: request
func (_m *Author) FindByName(request model.NameAuthorRequest) ([]model.Author, error) {
	ret := _m.Called(request)

	var r0 []model.Author
	if rf, ok := ret.Get(0).(func(model.NameAuthorRequest) []model.Author); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.NameAuthorRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserID provides a mock function with given fields: request
func (_m *Author) FindByUserID(request model.UserIDAuthorRequest) (*model.Author, error) {
	ret := _m.Called(request)

	var r0 *model.Author
	if rf, ok := ret.Get(0).(func(model.UserIDAuthorRequest) *model.Author); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.UserIDAuthorRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: request
func (_m *Author) Update(request model.UpdateAuthorRequest) (int, error) {
	ret := _m.Called(request)

	var r0 int
	if rf, ok := ret.Get(0).(func(model.UpdateAuthorRequest) int); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.UpdateAuthorRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
