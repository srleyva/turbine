// Code generated by MockGen. DO NOT EDIT.
// Source: proto/authentication/authentication.micro.go

// Package mocks_authentication is a generated GoMock package.
package mocks_authentication

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	client "github.com/micro/go-micro/client"
	authentication "github.com/srleyva/turbine/authentication-service/proto/authentication"
	user "github.com/srleyva/turbine/user-service/proto/user"
	reflect "reflect"
)

// MockAuthenticationService is a mock of AuthenticationService interface
type MockAuthenticationService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationServiceMockRecorder
}

// MockAuthenticationServiceMockRecorder is the mock recorder for MockAuthenticationService
type MockAuthenticationServiceMockRecorder struct {
	mock *MockAuthenticationService
}

// NewMockAuthenticationService creates a new mock instance
func NewMockAuthenticationService(ctrl *gomock.Controller) *MockAuthenticationService {
	mock := &MockAuthenticationService{ctrl: ctrl}
	mock.recorder = &MockAuthenticationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationService) EXPECT() *MockAuthenticationServiceMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockAuthenticationService) Login(ctx context.Context, in *authentication.LoginRequest, opts ...client.CallOption) (*authentication.LoginResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*authentication.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAuthenticationServiceMockRecorder) Login(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthenticationService)(nil).Login), varargs...)
}

// Register mocks base method
func (m *MockAuthenticationService) Register(ctx context.Context, in *user.User, opts ...client.CallOption) (*authentication.RegisterResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*authentication.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockAuthenticationServiceMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthenticationService)(nil).Register), varargs...)
}

// MockAuthenticationServiceHandler is a mock of AuthenticationServiceHandler interface
type MockAuthenticationServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationServiceHandlerMockRecorder
}

// MockAuthenticationServiceHandlerMockRecorder is the mock recorder for MockAuthenticationServiceHandler
type MockAuthenticationServiceHandlerMockRecorder struct {
	mock *MockAuthenticationServiceHandler
}

// NewMockAuthenticationServiceHandler creates a new mock instance
func NewMockAuthenticationServiceHandler(ctrl *gomock.Controller) *MockAuthenticationServiceHandler {
	mock := &MockAuthenticationServiceHandler{ctrl: ctrl}
	mock.recorder = &MockAuthenticationServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationServiceHandler) EXPECT() *MockAuthenticationServiceHandlerMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockAuthenticationServiceHandler) Login(arg0 context.Context, arg1 *authentication.LoginRequest, arg2 *authentication.LoginResponse) error {
	ret := m.ctrl.Call(m, "Login", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockAuthenticationServiceHandlerMockRecorder) Login(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthenticationServiceHandler)(nil).Login), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockAuthenticationServiceHandler) Register(arg0 context.Context, arg1 *user.User, arg2 *authentication.RegisterResponse) error {
	ret := m.ctrl.Call(m, "Register", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockAuthenticationServiceHandlerMockRecorder) Register(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthenticationServiceHandler)(nil).Register), arg0, arg1, arg2)
}
