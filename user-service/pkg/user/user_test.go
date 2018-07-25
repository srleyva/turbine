package user

import (
	proto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"reflect"
	"testing"
)

// Mocks
type mockGetUsersStream struct {
	Users []proto.User
}

func (s mockGetUsersStream) SendMsg(interface{}) error {
	return nil
}

func (s mockGetUsersStream) RecvMsg(interface{}) error {
	return nil
}

func (s mockGetUsersStream) Close() error {
	return nil
}

func (s mockGetUsersStream) Send(user *proto.User) error {
	updated := append(s.Users, *user)
	s.Users = updated
	return nil
}

// Local DataStore Tests
var datastore = DataStore{}
var service = Service{&datastore}
var testUser = &proto.User{
	FirstName: "test",
	LastName:  "user",
	Username:  "tuser",
	Password:  "Test",
}

func TestCreate(t *testing.T) {
	user, err := datastore.Create(testUser)
	if err != nil {
		t.Errorf("error returned where not expected: %v", err)
	}

	if !reflect.DeepEqual(testUser, user) {
		t.Errorf("Expected: %v \n Actual: %v", testUser, user)
	}

}

// Service Functions Tests
func TestCreateUser(t *testing.T) {
	response := new(proto.UserResponse)
	testUser := &proto.User{
		FirstName: "test",
		LastName:  "user",
		Username:  "tuser",
		Password:  "Test",
	}
	if err := service.CreateUser(context.TODO(), testUser, response); err != nil {
		t.Errorf("error returned where not expected: %v", err)
	}

	if response.User.Password != "" {
		t.Errorf("Password returned to client on create")
	}

	testUser.Username = ""
	if err := service.CreateUser(context.TODO(), testUser, response); err == nil {
		t.Errorf("error not returned where not expected: %v", testUser)
	}

}

func TestGetUsers(t *testing.T) {
	testStream := mockGetUsersStream{}

	t.Run("test that it returns all users error free", func(t *testing.T) {
		testReq := &proto.UsersRequest{All: true}
		err := service.GetUsers(context.TODO(), testReq, testStream)
		if err != nil {
			t.Errorf("err returned where not expected: %v", err)
		}

	})
}
