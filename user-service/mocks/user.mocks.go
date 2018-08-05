// Code generated by MockGen. DO NOT EDIT.
// Source: proto/user/user.micro.go

// Package mocks_user is a generated GoMock package.
package mocks_user

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	client "github.com/micro/go-micro/client"
	user "github.com/srleyva/turbine/user-service/proto/user"
	reflect "reflect"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserService) CreateUser(ctx context.Context, in *user.User, opts ...client.CallOption) (*user.UserResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUser", varargs...)
	ret0, _ := ret[0].(*user.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserServiceMockRecorder) CreateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), varargs...)
}

// GetUsers mocks base method
func (m *MockUserService) GetUsers(ctx context.Context, in *user.UsersRequest, opts ...client.CallOption) (user.UserService_GetUsersService, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsers", varargs...)
	ret0, _ := ret[0].(user.UserService_GetUsersService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockUserServiceMockRecorder) GetUsers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserService)(nil).GetUsers), varargs...)
}

// GetUser mocks base method
func (m *MockUserService) GetUser(ctx context.Context, in *user.UserRequest, opts ...client.CallOption) (*user.UserResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*user.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserServiceMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserService)(nil).GetUser), varargs...)
}

// MockUserService_GetUsersService is a mock of UserService_GetUsersService interface
type MockUserService_GetUsersService struct {
	ctrl     *gomock.Controller
	recorder *MockUserService_GetUsersServiceMockRecorder
}

// MockUserService_GetUsersServiceMockRecorder is the mock recorder for MockUserService_GetUsersService
type MockUserService_GetUsersServiceMockRecorder struct {
	mock *MockUserService_GetUsersService
}

// NewMockUserService_GetUsersService creates a new mock instance
func NewMockUserService_GetUsersService(ctrl *gomock.Controller) *MockUserService_GetUsersService {
	mock := &MockUserService_GetUsersService{ctrl: ctrl}
	mock.recorder = &MockUserService_GetUsersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService_GetUsersService) EXPECT() *MockUserService_GetUsersServiceMockRecorder {
	return m.recorder
}

// SendMsg mocks base method
func (m *MockUserService_GetUsersService) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockUserService_GetUsersServiceMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserService_GetUsersService)(nil).SendMsg), arg0)
}

// RecvMsg mocks base method
func (m *MockUserService_GetUsersService) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockUserService_GetUsersServiceMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserService_GetUsersService)(nil).RecvMsg), arg0)
}

// Close mocks base method
func (m *MockUserService_GetUsersService) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockUserService_GetUsersServiceMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUserService_GetUsersService)(nil).Close))
}

// Recv mocks base method
func (m *MockUserService_GetUsersService) Recv() (*user.User, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockUserService_GetUsersServiceMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockUserService_GetUsersService)(nil).Recv))
}

// MockUserServiceHandler is a mock of UserServiceHandler interface
type MockUserServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceHandlerMockRecorder
}

// MockUserServiceHandlerMockRecorder is the mock recorder for MockUserServiceHandler
type MockUserServiceHandlerMockRecorder struct {
	mock *MockUserServiceHandler
}

// NewMockUserServiceHandler creates a new mock instance
func NewMockUserServiceHandler(ctrl *gomock.Controller) *MockUserServiceHandler {
	mock := &MockUserServiceHandler{ctrl: ctrl}
	mock.recorder = &MockUserServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserServiceHandler) EXPECT() *MockUserServiceHandlerMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserServiceHandler) CreateUser(arg0 context.Context, arg1 *user.User, arg2 *user.UserResponse) error {
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserServiceHandlerMockRecorder) CreateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserServiceHandler)(nil).CreateUser), arg0, arg1, arg2)
}

// GetUsers mocks base method
func (m *MockUserServiceHandler) GetUsers(arg0 context.Context, arg1 *user.UsersRequest, arg2 user.UserService_GetUsersStream) error {
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockUserServiceHandlerMockRecorder) GetUsers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserServiceHandler)(nil).GetUsers), arg0, arg1, arg2)
}

// GetUser mocks base method
func (m *MockUserServiceHandler) GetUser(arg0 context.Context, arg1 *user.UserRequest, arg2 *user.UserResponse) error {
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserServiceHandlerMockRecorder) GetUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserServiceHandler)(nil).GetUser), arg0, arg1, arg2)
}

// MockUserService_GetUsersStream is a mock of UserService_GetUsersStream interface
type MockUserService_GetUsersStream struct {
	ctrl     *gomock.Controller
	recorder *MockUserService_GetUsersStreamMockRecorder
}

// MockUserService_GetUsersStreamMockRecorder is the mock recorder for MockUserService_GetUsersStream
type MockUserService_GetUsersStreamMockRecorder struct {
	mock *MockUserService_GetUsersStream
}

// NewMockUserService_GetUsersStream creates a new mock instance
func NewMockUserService_GetUsersStream(ctrl *gomock.Controller) *MockUserService_GetUsersStream {
	mock := &MockUserService_GetUsersStream{ctrl: ctrl}
	mock.recorder = &MockUserService_GetUsersStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService_GetUsersStream) EXPECT() *MockUserService_GetUsersStreamMockRecorder {
	return m.recorder
}

// SendMsg mocks base method
func (m *MockUserService_GetUsersStream) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockUserService_GetUsersStreamMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserService_GetUsersStream)(nil).SendMsg), arg0)
}

// RecvMsg mocks base method
func (m *MockUserService_GetUsersStream) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockUserService_GetUsersStreamMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserService_GetUsersStream)(nil).RecvMsg), arg0)
}

// Close mocks base method
func (m *MockUserService_GetUsersStream) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockUserService_GetUsersStreamMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUserService_GetUsersStream)(nil).Close))
}

// Send mocks base method
func (m *MockUserService_GetUsersStream) Send(arg0 *user.User) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockUserService_GetUsersStreamMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockUserService_GetUsersStream)(nil).Send), arg0)
}
