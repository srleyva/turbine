package Authentication

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	proto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	userProto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"time"
)

const (
	// Key (Should come from somewhere else). Vault?
	Key = "secret"
)

type Service struct {
	UserClient userProto.UserService
}

func (s *Service) Login(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("invalid Username or Password")
	}

	resp, err := s.UserClient.GetUser(context.Background(), &userProto.UserRequest{req.Username})
	if err != nil {
		return errors.New("invalid Username or Password")
	}

	if CheckPasswordHash(req.Password, resp.User.Password) {
		token, err := genToken(resp.User.UID)
		if err != nil {
			return err
		}
		res.Token = token
		res.Username = resp.User.Username
		res.FirstName = resp.User.FirstName
		res.LastName = resp.User.LastName
		return nil
	} else {
		return errors.New("invalid Username or Password")
	}

	return nil
}

func (s *Service) Register(ctx context.Context, req *userProto.User, res *proto.RegisterResponse) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("empty username or password")
	}

	_, err := s.UserClient.GetUser(context.Background(), &userProto.UserRequest{req.Username})
	if err == nil {
		return errors.New("username exists")
	}

	password, err := HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := userProto.User(*req)
	user.Password = password

	resp, err := s.UserClient.CreateUser(context.Background(), &user)
	if err != nil {
		return err
	}
	res.Created = true
	res.User = resp.User
	return nil
}

func genToken(uid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	return token.SignedString([]byte(Key))
}
