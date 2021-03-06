// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

/*
Package go_micro_srv_user is a generated protocol buffer package.

It is generated from these files:
	proto/user/user.proto

It has these top-level messages:
	User
	UserRequest
	UsersRequest
	UserResponse
*/
package go_micro_srv_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error)
	GetUsers(ctx context.Context, in *UsersRequest, opts ...client.CallOption) (UserService_GetUsersService, error)
	GetUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) CreateUser(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.CreateUser", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUsers(ctx context.Context, in *UsersRequest, opts ...client.CallOption) (UserService_GetUsersService, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUsers", &UsersRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &userServiceGetUsers{stream}, nil
}

type UserService_GetUsersService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*User, error)
}

type userServiceGetUsers struct {
	stream client.Stream
}

func (x *userServiceGetUsers) Close() error {
	return x.stream.Close()
}

func (x *userServiceGetUsers) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceGetUsers) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceGetUsers) Recv() (*User, error) {
	m := new(User)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userService) GetUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUser", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	CreateUser(context.Context, *User, *UserResponse) error
	GetUsers(context.Context, *UsersRequest, UserService_GetUsersStream) error
	GetUser(context.Context, *UserRequest, *UserResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	type userService interface {
		CreateUser(ctx context.Context, in *User, out *UserResponse) error
		GetUsers(ctx context.Context, stream server.Stream) error
		GetUser(ctx context.Context, in *UserRequest, out *UserResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) CreateUser(ctx context.Context, in *User, out *UserResponse) error {
	return h.UserServiceHandler.CreateUser(ctx, in, out)
}

func (h *userServiceHandler) GetUsers(ctx context.Context, stream server.Stream) error {
	m := new(UsersRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UserServiceHandler.GetUsers(ctx, m, &userServiceGetUsersStream{stream})
}

type UserService_GetUsersStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*User) error
}

type userServiceGetUsersStream struct {
	stream server.Stream
}

func (x *userServiceGetUsersStream) Close() error {
	return x.stream.Close()
}

func (x *userServiceGetUsersStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceGetUsersStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceGetUsersStream) Send(m *User) error {
	return x.stream.Send(m)
}

func (h *userServiceHandler) GetUser(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserServiceHandler.GetUser(ctx, in, out)
}
