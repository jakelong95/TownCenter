package mocks

import helpers "github.com/jakelong95/TownCenter/helpers"
import mock "github.com/stretchr/testify/mock"
import models "github.com/jakelong95/TownCenter/models"
import multipart "mime/multipart"
import uuid "github.com/pborman/uuid"

// RoasterI is an autogenerated mock type for the RoasterI type
type RoasterI struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: id
func (_m *RoasterI) CreateAccount(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *RoasterI) Delete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: _a0, _a1
func (_m *RoasterI) GetAll(_a0 int, _a1 int) ([]*models.Roaster, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*models.Roaster
	if rf, ok := ret.Get(0).(func(int, int) []*models.Roaster); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Roaster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0
func (_m *RoasterI) GetByID(_a0 string) (*models.Roaster, error) {
	ret := _m.Called(_a0)

	var r0 *models.Roaster
	if rf, ok := ret.Get(0).(func(string) *models.Roaster); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Roaster)
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
func (_m *RoasterI) Insert(_a0 *models.Roaster) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Roaster) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Profile provides a mock function with given fields: _a0, _a1, _a2
func (_m *RoasterI) Profile(_a0 string, _a1 string, _a2 multipart.File) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, multipart.File) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *RoasterI) Update(_a0 *models.Roaster, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Roaster, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ helpers.RoasterI = (*RoasterI)(nil)
