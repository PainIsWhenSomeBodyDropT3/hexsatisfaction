// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mock

import (
	model "github.com/JesusG2000/hexsatisfaction/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// UserRole is an autogenerated mock type for the UserRole type
type UserRole struct {
	mock.Mock
}

// FindAllUser provides a mock function with given fields:
func (_m *UserRole) FindAllUser() ([]model.User, error) {
	ret := _m.Called()

	var r0 []model.User
	if rf, ok := ret.Get(0).(func() []model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
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