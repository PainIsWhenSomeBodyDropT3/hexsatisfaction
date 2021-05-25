// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mock

import (
	model "github.com/JesusG2000/hexsatisfaction/internal/model"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Purchase is an autogenerated mock type for the Purchase type
type Purchase struct {
	mock.Mock
}

// Create provides a mock function with given fields: purchase
func (_m *Purchase) Create(purchase model.Purchase) (int, error) {
	ret := _m.Called(purchase)

	var r0 int
	if rf, ok := ret.Get(0).(func(model.Purchase) int); ok {
		r0 = rf(purchase)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Purchase) error); ok {
		r1 = rf(purchase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Purchase) Delete(id int) (int, error) {
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

// DeleteByFileID provides a mock function with given fields: id
func (_m *Purchase) DeleteByFileID(id int) (int, error) {
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

// FindAfterDate provides a mock function with given fields: start
func (_m *Purchase) FindAfterDate(start time.Time) ([]model.Purchase, error) {
	ret := _m.Called(start)

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func(time.Time) []model.Purchase); ok {
		r0 = rf(start)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time) error); ok {
		r1 = rf(start)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *Purchase) FindAll() ([]model.Purchase, error) {
	ret := _m.Called()

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func() []model.Purchase); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
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
func (_m *Purchase) FindAllByUserID(id int) ([]model.Purchase, error) {
	ret := _m.Called(id)

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func(int) []model.Purchase); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
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

// FindBeforeDate provides a mock function with given fields: end
func (_m *Purchase) FindBeforeDate(end time.Time) ([]model.Purchase, error) {
	ret := _m.Called(end)

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func(time.Time) []model.Purchase); ok {
		r0 = rf(end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time) error); ok {
		r1 = rf(end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByFileID provides a mock function with given fields: id
func (_m *Purchase) FindByFileID(id int) ([]model.Purchase, error) {
	ret := _m.Called(id)

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func(int) []model.Purchase); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
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
func (_m *Purchase) FindByID(id int) (*model.Purchase, error) {
	ret := _m.Called(id)

	var r0 *model.Purchase
	if rf, ok := ret.Get(0).(func(int) *model.Purchase); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Purchase)
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
func (_m *Purchase) FindByPeriod(start time.Time, end time.Time) ([]model.Purchase, error) {
	ret := _m.Called(start, end)

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) []model.Purchase); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
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

// FindByUserIDAfterDate provides a mock function with given fields: id, start
func (_m *Purchase) FindByUserIDAfterDate(id int, start time.Time) ([]model.Purchase, error) {
	ret := _m.Called(id, start)

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func(int, time.Time) []model.Purchase); ok {
		r0 = rf(id, start)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, time.Time) error); ok {
		r1 = rf(id, start)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserIDAndFileID provides a mock function with given fields: userID, fileID
func (_m *Purchase) FindByUserIDAndFileID(userID int, fileID int) ([]model.Purchase, error) {
	ret := _m.Called(userID, fileID)

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func(int, int) []model.Purchase); ok {
		r0 = rf(userID, fileID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userID, fileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserIDAndPeriod provides a mock function with given fields: id, start, end
func (_m *Purchase) FindByUserIDAndPeriod(id int, start time.Time, end time.Time) ([]model.Purchase, error) {
	ret := _m.Called(id, start, end)

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func(int, time.Time, time.Time) []model.Purchase); ok {
		r0 = rf(id, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, time.Time, time.Time) error); ok {
		r1 = rf(id, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserIDBeforeDate provides a mock function with given fields: id, end
func (_m *Purchase) FindByUserIDBeforeDate(id int, end time.Time) ([]model.Purchase, error) {
	ret := _m.Called(id, end)

	var r0 []model.Purchase
	if rf, ok := ret.Get(0).(func(int, time.Time) []model.Purchase); ok {
		r0 = rf(id, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Purchase)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, time.Time) error); ok {
		r1 = rf(id, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLast provides a mock function with given fields:
func (_m *Purchase) FindLast() (*model.Purchase, error) {
	ret := _m.Called()

	var r0 *model.Purchase
	if rf, ok := ret.Get(0).(func() *model.Purchase); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Purchase)
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

// FindLastByUserID provides a mock function with given fields: id
func (_m *Purchase) FindLastByUserID(id int) (*model.Purchase, error) {
	ret := _m.Called(id)

	var r0 *model.Purchase
	if rf, ok := ret.Get(0).(func(int) *model.Purchase); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Purchase)
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
