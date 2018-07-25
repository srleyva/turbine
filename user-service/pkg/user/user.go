package user

import (
	"errors"
	"github.com/satori/go.uuid"
	proto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
)

type UserStore interface {
	Create(*proto.User) (*proto.User, error)
	GetAll() (*[]proto.User, error)
	Get(string) (*proto.User, error)
}

type Service struct {
	Users UserStore
}

func (s *Service) CreateUser(ctx context.Context, req *proto.User, res *proto.UserResponse) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("empty username or password in request")
	}
	req.UID = uuid.NewV4().String()
	response, err := s.Users.Create(req)
	if err != nil {
		return err
	}
	user := *response
	user.Password = ""
	res.User = &user
	return nil
}

func (s *Service) GetUsers(ctx context.Context, req *proto.UsersRequest, stream proto.UserService_GetUsersStream) error {
	allUsers, _ := s.Users.GetAll()
	for index, user := range *allUsers {
		user.Password = ""
		if req.All || req.Count > int64(index) {
			if err := stream.Send(&user); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Service) GetUser(ctx context.Context, req *proto.UserRequest, res *proto.UserResponse) error {
	user, err := s.Users.Get(req.Username)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}
