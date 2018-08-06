package authentication_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/srleyva/turbine/authentication-service/pkg/authentication"
	proto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	userMocks "github.com/srleyva/turbine/user-service/mocks"
	userProto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"testing"
)

type tester struct {
	ctrl        *gomock.Controller
	userService *userMocks.MockUserService
	service     *Service
}

func newTester(t *testing.T) *tester {
	ctrl := gomock.NewController(t)
	mockUser := userMocks.NewMockUserService(ctrl)
	service := &Service{mockUser}
	return &tester{ctrl, mockUser, service}
}

func TestLogin(t *testing.T) {
	authTester := newTester(t)
	defer authTester.ctrl.Finish()

	// As the password is hashed in database
	testPassword, _ := HashPassword("test")

	authTester.userService.EXPECT().GetUser(
		context.Background(),
		&userProto.UserRequest{Username: "thisdude"}).Return(
		&userProto.UserResponse{&userProto.User{Username: "thisdude", Password: testPassword}}, nil)

	req := &proto.LoginRequest{Username: "thisdude", Password: "test"}
	res := &proto.LoginResponse{}

	if err := authTester.service.Login(context.Background(), req, res); err != nil {
		t.Errorf("error returned where not expected: %v", err)
	}
}

func TestRegister(t *testing.T) {
	authTester := newTester(t)
	defer authTester.ctrl.Finish()

	testUser := &userProto.User{UID: "this-is-an-id", Username: "thisdude"}

	authTester.userService.EXPECT().GetUser(
		context.Background(),
		&userProto.UserRequest{Username: "thisdude"}).Return(
		&userProto.UserResponse{nil}, errors.New("user does not exists"))

	authTester.userService.EXPECT().CreateUser(
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).Return(&userProto.UserResponse{User: testUser}, nil)

	req := &userProto.User{Username: "thisdude", Password: "password"}
	res := &proto.RegisterResponse{}

	if err := authTester.service.Register(context.Background(), req, res); err != nil {
		t.Errorf("error returned where not expected: %v", err)
	}

	if res.User != testUser {
		t.Errorf("incorrect information returned \n Have: %v \n Want: %v", res.User, testUser)
	}
}
