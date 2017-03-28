package mocks

import helpers "github.com/jakelong95/TownCenter/helpers"
import mock "github.com/stretchr/testify/mock"
import models "github.com/jakelong95/TownCenter/models"

// ResetI is an autogenerated mock type for the ResetI type
type ResetI struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *ResetI) Get(_a0 string) (*models.Token, error) {
	ret := _m.Called(_a0)

	var r0 *models.Token
	if rf, ok := ret.Get(0).(func(string) *models.Token); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *ResetI) Insert(_a0 *models.Token) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Token) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetStatus provides a mock function with given fields: _a0, _a1
func (_m *ResetI) SetStatus(_a0 *models.Token, _a1 models.TokenStatus) (*models.Token, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.Token
	if rf, ok := ret.Get(0).(func(*models.Token, models.TokenStatus) *models.Token); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Token, models.TokenStatus) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ helpers.ResetI = (*ResetI)(nil)
