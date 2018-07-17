package user

import (
	proto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"reflect"
	"testing"
)

// Mocks
type mockGetUsersStream struct {
	Users []string
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
	updated := append(s.Users, user.Username)
	s.Users = updated
	return nil
}

// Vars
var response = &proto.Response{Create: false}

// Local DataStore Tests
var datastore = DataStore{}
var service = Service{&datastore}

func TestCreate(t *testing.T) {
	testUser := &proto.User{
		FirstName: "test",
		LastName:  "user",
		Username:  "tuser",
		Password:  "Test",
	}
	user, err := datastore.Create(testUser)
	if err != nil {
		t.Errorf("error returned where not expected: %v", err)
	}

	if user.GetPassword() != "" {
		t.Errorf("password returned to client")
	}

	if !reflect.DeepEqual(testUser, user) {
		t.Errorf("Expected: %v \n Actual: %v", testUser, user)
	}

}

// Service Functions Tests
// TODO: Tests need to return more than just no error
func TestCreateUser(t *testing.T) {
	testUser := &proto.User{
		FirstName: "test",
		LastName:  "user",
		Username:  "tuser",
		Password:  "Test",
	}
	if err := service.CreateUser(context.TODO(), testUser, response); err != nil {
		t.Errorf("error returned where not expected: %v", err)
	}

	if !response.Create || !reflect.DeepEqual(testUser, response.User) {
		t.Errorf("user not create correctly %v", response)
	}

	testUser.Username = ""
	if err := service.CreateUser(context.TODO(), testUser, response); err == nil {
		t.Errorf("error not returned where not expected: %v", testUser)
	}

}

func TestGetUsers(t *testing.T) {
	t.Run("test that it returns all users error free", func(t *testing.T) {
		testReq := &proto.UsersRequest{true, 100}
		testStream := mockGetUsersStream{}
		err := service.GetUsers(context.TODO(), testReq, testStream)
		if err != nil {
			t.Errorf("err returned where not expected: %v", err)
		}
	})
}
