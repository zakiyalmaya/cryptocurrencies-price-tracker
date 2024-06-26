// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

// MockTrackerService is a mock of TrackerService interface.
type MockTrackerService struct {
	ctrl     *gomock.Controller
	recorder *MockTrackerServiceMockRecorder
}

// MockTrackerServiceMockRecorder is the mock recorder for MockTrackerService.
type MockTrackerServiceMockRecorder struct {
	mock *MockTrackerService
}

// NewMockTrackerService creates a new mock instance.
func NewMockTrackerService(ctrl *gomock.Controller) *MockTrackerService {
	mock := &MockTrackerService{ctrl: ctrl}
	mock.recorder = &MockTrackerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackerService) EXPECT() *MockTrackerServiceMockRecorder {
	return m.recorder
}

// AddUserTrackedCoin mocks base method.
func (m *MockTrackerService) AddUserTrackedCoin(req *model.AddUserTrackedCoinRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserTrackedCoin", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserTrackedCoin indicates an expected call of AddUserTrackedCoin.
func (mr *MockTrackerServiceMockRecorder) AddUserTrackedCoin(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserTrackedCoin", reflect.TypeOf((*MockTrackerService)(nil).AddUserTrackedCoin), req)
}

// DeleteUserTrackedCoin mocks base method.
func (m *MockTrackerService) DeleteUserTrackedCoin(userID int, coinID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserTrackedCoin", userID, coinID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserTrackedCoin indicates an expected call of DeleteUserTrackedCoin.
func (mr *MockTrackerServiceMockRecorder) DeleteUserTrackedCoin(userID, coinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserTrackedCoin", reflect.TypeOf((*MockTrackerService)(nil).DeleteUserTrackedCoin), userID, coinID)
}

// GetAssetList mocks base method.
func (m *MockTrackerService) GetAssetList(req *model.AssetRequest) (*model.AssetsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetList", req)
	ret0, _ := ret[0].(*model.AssetsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetList indicates an expected call of GetAssetList.
func (mr *MockTrackerServiceMockRecorder) GetAssetList(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetList", reflect.TypeOf((*MockTrackerService)(nil).GetAssetList), req)
}

// GetUserTrackedList mocks base method.
func (m *MockTrackerService) GetUserTrackedList(username string) (*model.UserTrackedCoin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserTrackedList", username)
	ret0, _ := ret[0].(*model.UserTrackedCoin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserTrackedList indicates an expected call of GetUserTrackedList.
func (mr *MockTrackerServiceMockRecorder) GetUserTrackedList(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserTrackedList", reflect.TypeOf((*MockTrackerService)(nil).GetUserTrackedList), username)
}
