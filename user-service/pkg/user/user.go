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

	if _, err := s.Users.Get(req.Username); err == nil {
		return errors.New("Username already exists")
	}
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	req.UID = uid.String()
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
	users := *allUsers
	for i := 0; i < len(users); i++ {
		user := users[i]
		if user != (proto.User{}) {
			if req.All || req.Count > int64(i) {
				if err := stream.Send(&user); err != nil {
					return err
				}
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
