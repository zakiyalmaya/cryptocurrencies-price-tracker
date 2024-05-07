// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

// MockTrackerRepository is a mock of TrackerRepository interface.
type MockTrackerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTrackerRepositoryMockRecorder
}

// MockTrackerRepositoryMockRecorder is the mock recorder for MockTrackerRepository.
type MockTrackerRepositoryMockRecorder struct {
	mock *MockTrackerRepository
}

// NewMockTrackerRepository creates a new mock instance.
func NewMockTrackerRepository(ctrl *gomock.Controller) *MockTrackerRepository {
	mock := &MockTrackerRepository{ctrl: ctrl}
	mock.recorder = &MockTrackerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackerRepository) EXPECT() *MockTrackerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTrackerRepository) Create(req *model.TrackerEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTrackerRepositoryMockRecorder) Create(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTrackerRepository)(nil).Create), req)
}

// Delete mocks base method.
func (m *MockTrackerRepository) Delete(userID int, coinID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userID, coinID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTrackerRepositoryMockRecorder) Delete(userID, coinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrackerRepository)(nil).Delete), userID, coinID)
}

// GetByUserIDAndCoinID mocks base method.
func (m *MockTrackerRepository) GetByUserIDAndCoinID(userID int, coinID string) (*model.TrackerEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserIDAndCoinID", userID, coinID)
	ret0, _ := ret[0].(*model.TrackerEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserIDAndCoinID indicates an expected call of GetByUserIDAndCoinID.
func (mr *MockTrackerRepositoryMockRecorder) GetByUserIDAndCoinID(userID, coinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserIDAndCoinID", reflect.TypeOf((*MockTrackerRepository)(nil).GetByUserIDAndCoinID), userID, coinID)
}

// GetByUsername mocks base method.
func (m *MockTrackerRepository) GetByUsername(username string) (*model.UserTrackedCoin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(*model.UserTrackedCoin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockTrackerRepositoryMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockTrackerRepository)(nil).GetByUsername), username)
}