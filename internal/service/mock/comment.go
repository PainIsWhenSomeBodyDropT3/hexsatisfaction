// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mock

import (
	model "github.com/JesusG2000/hexsatisfaction/internal/model"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Comment is an autogenerated mock type for the Comment type
type Comment struct {
	mock.Mock
}

// Create provides a mock function with given fields: comment
func (_m *Comment) Create(comment model.Comment) (int, error) {
	ret := _m.Called(comment)

	var r0 int
	if rf, ok := ret.Get(0).(func(model.Comment) int); ok {
		r0 = rf(comment)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Comment) error); ok {
		r1 = rf(comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Comment) Delete(id int) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *Comment) FindAll() ([]model.Comment, error) {
	ret := _m.Called()

	var r0 []model.Comment
	if rf, ok := ret.Get(0).(func() []model.Comment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Comment)
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

// FindAllByUserID provides a mock function with given fields: id
func (_m *Comment) FindAllByUserID(id int) ([]model.Comment, error) {
	ret := _m.Called(id)

	var r0 []model.Comment
	if rf, ok := ret.Get(0).(func(int) []model.Comment); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *Comment) FindByID(id int) (*model.Comment, error) {
	ret := _m.Called(id)

	var r0 *model.Comment
	if rf, ok := ret.Get(0).(func(int) *model.Comment); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByPeriod provides a mock function with given fields: start, end
func (_m *Comment) FindByPeriod(start time.Time, end time.Time) ([]model.Comment, error) {
	ret := _m.Called(start, end)

	var r0 []model.Comment
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) []model.Comment); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time, time.Time) error); ok {
		r1 = rf(start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByPurchaseID provides a mock function with given fields: id
func (_m *Comment) FindByPurchaseID(id int) ([]model.Comment, error) {
	ret := _m.Called(id)

	var r0 []model.Comment
	if rf, ok := ret.Get(0).(func(int) []model.Comment); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByText provides a mock function with given fields: text
func (_m *Comment) FindByText(text string) ([]model.Comment, error) {
	ret := _m.Called(text)

	var r0 []model.Comment
	if rf, ok := ret.Get(0).(func(string) []model.Comment); ok {
		r0 = rf(text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserIDAndPurchaseID provides a mock function with given fields: userID, purchaseID
func (_m *Comment) FindByUserIDAndPurchaseID(userID int, purchaseID int) ([]model.Comment, error) {
	ret := _m.Called(userID, purchaseID)

	var r0 []model.Comment
	if rf, ok := ret.Get(0).(func(int, int) []model.Comment); ok {
		r0 = rf(userID, purchaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userID, purchaseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, comment
func (_m *Comment) Update(id int, comment model.Comment) (int, error) {
	ret := _m.Called(id, comment)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, model.Comment) int); ok {
		r0 = rf(id, comment)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, model.Comment) error); ok {
		r1 = rf(id, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}