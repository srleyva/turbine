package Authentication

import (
	"errors"
	proto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	userProto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
)

type Service struct {
	UserClient userProto.UserService
}

func (s *Service) Login(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("invalid Username or Password")
	}

	resp, err := s.UserClient.GetUser(context.Background(), &userProto.UserRequest{Username: req.Username})
	if err != nil {
		return errors.New("invalid Username or Password")
	}

	if CheckPasswordHash(req.Password, resp.User.Password) {
		res.Token = "You've got a token"
		return nil
	} else {
		return errors.New("invalid Username or Password")
	}

	return nil
}

func (s *Service) Register(ctx context.Context, req *proto.RegisterRequest, res *proto.RegisterResponse) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("empty username or password")
	}

	password, err := HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := userProto.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Password:  password,
	}

	resp, err := s.UserClient.CreateUser(context.Background(), &user)
	if err != nil {
		return err
	}
	res.Username = resp.User.Username
	res.UID = resp.User.UID
	res.Created = true
	return nil
}
