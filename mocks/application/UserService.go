// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockUserService) Login(auth *model.AuthRequest) (*model.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", auth)
	ret0, _ := ret[0].(*model.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(auth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), auth)
}

// Logout mocks base method.
func (m *MockUserService) Logout(username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", username)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockUserServiceMockRecorder) Logout(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUserService)(nil).Logout), username)
}

// Register mocks base method.
func (m *MockUserService) Register(user *model.UserEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockUserServiceMockRecorder) Register(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserService)(nil).Register), user)
}
