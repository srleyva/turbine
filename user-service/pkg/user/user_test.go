package user_test

import (
	"github.com/golang/mock/gomock"
	userMock "github.com/srleyva/turbine/user-service/mocks"
	user "github.com/srleyva/turbine/user-service/pkg/user"
	proto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"reflect"
	"testing"
)

type tester struct {
	ctrl      *gomock.Controller
	datastore *user.DataStore
	service   *user.Service
}

func newTester(t *testing.T) *tester {
	ctrl := gomock.NewController(t)
	datastore := &user.DataStore{}
	service := &user.Service{datastore}
	return &tester{ctrl, datastore, service}
}

func TestCreateUser(t *testing.T) {
	createTester := newTester(t)

	testUser := &proto.User{Username: "thisdude", Password: "test"}
	resp := &proto.UserResponse{}

	if err := createTester.service.CreateUser(context.Background(), testUser, resp); err != nil {
		t.Errorf("error returned where not expected: %v", err)
	}

	testUser.Password = ""

	if !reflect.DeepEqual(testUser, resp.User) {
		t.Errorf("incorrect user information returned: \n Have: %v \n Want: %v", resp.User, testUser)
	}
}

func TestGetUsers(t *testing.T) {
	getUsersTester := newTester(t)

	// Set up users in database to retrieve
	testUser1 := &proto.User{Username: "thisdude", Password: "test"}
	testUser2 := &proto.User{Username: "thatdude", Password: "test"}
	testUser3 := &proto.User{Username: "thatotherdude", Password: "test"}
	testUser4 := &proto.User{Username: "dude", Password: "test"}
	users := []*proto.User{testUser1, testUser2, testUser3, testUser4}
	for _, user := range users {
		resp := &proto.UserResponse{}
		if err := getUsersTester.service.CreateUser(context.Background(), user, resp); err != nil {
			t.Errorf("err inserting user: %v", err)
		}
	}

	t.Run("test return all users", func(t *testing.T) {
		req := &proto.UsersRequest{All: true}
		stream := userMock.NewMockUserService_GetUsersStream(getUsersTester.ctrl)
		first := stream.EXPECT().Send(testUser1).Return(nil)
		second := stream.EXPECT().Send(testUser2).After(first).Return(nil)
		third := stream.EXPECT().Send(testUser3).After(second).Return(nil)
		stream.EXPECT().Send(testUser4).After(third).Return(nil)

		if err := getUsersTester.service.GetUsers(context.Background(), req, stream); err != nil {
			t.Errorf("err returned where not expected: %v", err)
		}
	})

	t.Run("return number users", func(t *testing.T) {
		req := &proto.UsersRequest{Count: 2}
		stream := userMock.NewMockUserService_GetUsersStream(getUsersTester.ctrl)
		first := stream.EXPECT().Send(testUser1).Return(nil)
		stream.EXPECT().Send(testUser2).After(first).Return(nil)

		if err := getUsersTester.service.GetUsers(context.Background(), req, stream); err != nil {
			t.Errorf("err returned where not expected: %v", err)
		}

	})
}

func TestGetUser(t *testing.T) {
	getUserTester := newTester(t)

	// Set up user in database to retrieve
	testUser := &proto.User{Username: "thisdude", Password: "test"}
	resp := &proto.UserResponse{}
	getUserTester.service.CreateUser(context.Background(), testUser, resp)

	//Test retrieval of user
	req := &proto.UserRequest{Username: "thisdude"}

	if err := getUserTester.service.GetUser(context.Background(), req, resp); err != nil {
		t.Errorf("err returned where not expected: %v", err)
	}

	if !reflect.DeepEqual(resp.User, testUser) {
		t.Errorf("incorrect user information returned: \n Have: %v \n Want: %v", resp.User, testUser)
	}
}
