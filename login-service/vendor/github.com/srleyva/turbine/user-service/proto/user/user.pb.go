// Code generated by protoc-gen-go. DO NOT EDIT.
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

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	UID       string `protobuf:"bytes,5,opt,name=UID" json:"UID,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

type UserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UsersRequest struct {
	All   bool  `protobuf:"varint,1,opt,name=all" json:"all,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *UsersRequest) Reset()                    { *m = UsersRequest{} }
func (m *UsersRequest) String() string            { return proto.CompactTextString(m) }
func (*UsersRequest) ProtoMessage()               {}
func (*UsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UsersRequest) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func (m *UsersRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type UserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*UserRequest)(nil), "go.micro.srv.user.UserRequest")
	proto.RegisterType((*UsersRequest)(nil), "go.micro.srv.user.UsersRequest")
	proto.RegisterType((*UserResponse)(nil), "go.micro.srv.user.UserResponse")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x37, 0x76, 0x57, 0xbb, 0xb3, 0x1e, 0xd6, 0xa0, 0x58, 0x2a, 0xfe, 0x21, 0x27, 0x45,
	0x88, 0xb2, 0x82, 0x17, 0x8f, 0x0a, 0xae, 0x20, 0x1e, 0x2a, 0x7b, 0x96, 0x58, 0x47, 0x29, 0xb4,
	0x4d, 0x4d, 0xd2, 0xf5, 0x25, 0x7c, 0x57, 0x5f, 0x41, 0x32, 0xd9, 0x1e, 0x16, 0x29, 0x5e, 0x4a,
	0x66, 0x7e, 0xdf, 0x7c, 0xf9, 0x32, 0x14, 0xf6, 0x1a, 0xa3, 0x9d, 0xbe, 0x68, 0x2d, 0x1a, 0xfa,
	0x48, 0xaa, 0xf9, 0xce, 0x87, 0x96, 0x55, 0x91, 0x1b, 0x2d, 0xad, 0x59, 0x4a, 0x0f, 0xc4, 0x37,
	0x83, 0xe1, 0xc2, 0xa2, 0xe1, 0x87, 0x00, 0xef, 0x85, 0xb1, 0xee, 0xa5, 0x56, 0x15, 0x26, 0xec,
	0x84, 0x9d, 0x8e, 0xb3, 0x31, 0x75, 0x9e, 0x54, 0x85, 0xfc, 0x00, 0xc6, 0xa5, 0xea, 0xe8, 0x06,
	0xd1, 0xd8, 0x37, 0x08, 0xa6, 0x10, 0x7b, 0x33, 0x62, 0x51, 0x60, 0x5d, 0xed, 0x59, 0xa3, 0xac,
	0xfd, 0xd2, 0xe6, 0x2d, 0x19, 0x06, 0xd6, 0xd5, 0x7c, 0x0a, 0xd1, 0xe2, 0xe1, 0x2e, 0x19, 0x51,
	0xdb, 0x1f, 0xc5, 0x19, 0x4c, 0x7c, 0x9a, 0x0c, 0x3f, 0x5b, 0xb4, 0x6e, 0xcd, 0x98, 0xad, 0x1b,
	0x8b, 0x6b, 0xd8, 0xf6, 0x52, 0xdb, 0x69, 0xa7, 0x10, 0xa9, 0xb2, 0x24, 0x59, 0x9c, 0xf9, 0x23,
	0xdf, 0x85, 0x51, 0xae, 0xdb, 0xda, 0x51, 0xde, 0x28, 0x0b, 0x85, 0xb8, 0x09, 0x73, 0x19, 0xda,
	0x46, 0xd7, 0x16, 0xf9, 0x39, 0x0c, 0xbd, 0x27, 0x0d, 0x4e, 0x66, 0xfb, 0xf2, 0xcf, 0x8e, 0x24,
	0xc9, 0x49, 0x34, 0xfb, 0x61, 0x21, 0xe0, 0x33, 0x9a, 0x65, 0x91, 0x23, 0x9f, 0x03, 0xdc, 0x1a,
	0x54, 0x0e, 0x69, 0x87, 0x7d, 0xc3, 0xe9, 0x71, 0x9f, 0xeb, 0x2a, 0x84, 0x18, 0xf0, 0x39, 0xc4,
	0xf7, 0xe8, 0xe8, 0x45, 0xbc, 0x4f, 0xde, 0xbd, 0x35, 0xed, 0xbb, 0x48, 0x0c, 0x2e, 0x19, 0x7f,
	0x84, 0xad, 0x95, 0x13, 0x3f, 0xea, 0xbd, 0x37, 0xf8, 0xfc, 0x9f, 0xeb, 0x75, 0x93, 0x7e, 0x9d,
	0xab, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xe6, 0xaf, 0x6c, 0x53, 0x02, 0x00, 0x00,
}
