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

	if resp.User.Password == req.Password {
		res.Token = "You've got a token"
		return nil
	} else {
		return errors.New("invalid Username or Password")
	}

	return nil
}
