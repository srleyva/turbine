// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/authentication/authentication.proto

/*
Package go_micro_srv_authentication is a generated protocol buffer package.

It is generated from these files:
	proto/authentication/authentication.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	RegisterResponse
*/
package go_micro_srv_authentication

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_srv_user "github.com/srleyva/turbine/user-service/proto/user"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = go_micro_srv_user.User{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthenticationService service

type AuthenticationService interface {
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *go_micro_srv_user.User, opts ...client.CallOption) (*RegisterResponse, error)
}

type authenticationService struct {
	c    client.Client
	name string
}

func NewAuthenticationService(name string, c client.Client) AuthenticationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.authentication"
	}
	return &authenticationService{
		c:    c,
		name: name,
	}
}

func (c *authenticationService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "AuthenticationService.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationService) Register(ctx context.Context, in *go_micro_srv_user.User, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "AuthenticationService.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthenticationService service

type AuthenticationServiceHandler interface {
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Register(context.Context, *go_micro_srv_user.User, *RegisterResponse) error
}

func RegisterAuthenticationServiceHandler(s server.Server, hdlr AuthenticationServiceHandler, opts ...server.HandlerOption) {
	type authenticationService interface {
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Register(ctx context.Context, in *go_micro_srv_user.User, out *RegisterResponse) error
	}
	type AuthenticationService struct {
		authenticationService
	}
	h := &authenticationServiceHandler{hdlr}
	s.Handle(s.NewHandler(&AuthenticationService{h}, opts...))
}

type authenticationServiceHandler struct {
	AuthenticationServiceHandler
}

func (h *authenticationServiceHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AuthenticationServiceHandler.Login(ctx, in, out)
}

func (h *authenticationServiceHandler) Register(ctx context.Context, in *go_micro_srv_user.User, out *RegisterResponse) error {
	return h.AuthenticationServiceHandler.Register(ctx, in, out)
}
