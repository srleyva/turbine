package user

import (
	"errors"
	proto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
)

type UserStore interface {
	Create(*proto.User) (*proto.User, error)
	GetAll() ([]*proto.User, error)
}

// TODO: Implement Real DB backend This is for unit tests and local testing
type DataStore struct {
	users []*proto.User
}

func (d *DataStore) Create(user *proto.User) (*proto.User, error) {
	update := append(d.users, user)
	d.users = update
	user.Password = "" // Don't return password hash
	return user, nil
}

func (d *DataStore) GetAll() ([]*proto.User, error) {
	return d.users, nil
}

type Service struct {
	Users UserStore
}

func (s *Service) CreateUser(ctx context.Context, req *proto.User, res *proto.Response) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("empty username or password in request")
	}
	user, _ := s.Users.Create(req)
	res.User = user
	res.Create = true
	return nil
}

func (s *Service) GetUsers(ctx context.Context, req *proto.UsersRequest, stream proto.UserService_GetUsersStream) error {
	users, _ := s.Users.GetAll()
	for index, user := range users {
		if req.All || req.Count > int64(index) {
			user.Password = ""
			if err := stream.Send(user); err != nil {
				return err
			}
		}
	}
	return nil
}
