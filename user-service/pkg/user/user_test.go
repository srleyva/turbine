package user

import (
	proto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"reflect"
	"testing"
)

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

func TestCreateUser(t *testing.T) {
	testUser := &proto.User{
		FirstName: "test",
		LastName:  "user",
		Username:  "tuser",
		Password:  "Test",
	}

	var response = &proto.Response{Create: false}

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
