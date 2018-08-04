package login

import (
	"github.com/labstack/echo"
	authProto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	userProto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"io"
	"net/http"
)

const (
	// TODO Store in Vault
	Key = "secret"
)

type Handler struct {
	Auth authProto.AuthenticationService
	User userProto.UserService
}

func (h *Handler) GetAllUsers(c echo.Context) (err error) {
	users, err := getAllUsers(h.User)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) Login(c echo.Context) (err error) {
	u := &userProto.User{}
	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}

	response, err := login(h.Auth, u.Username, u.Password)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Register(c echo.Context) (err error) {
	u := &userProto.User{}
	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
	}

	response, err := register(h.Auth, u)
	if err != nil {
		if err.Error() == "empty username or password" {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
		}
		return err
	}

	return c.JSON(http.StatusCreated, response)

}

func getAllUsers(u userProto.UserService) (*[]userProto.User, error) {
	var users []userProto.User
	stream, err := u.GetUsers(context.Background(), &userProto.UsersRequest{All: true})
	if err != nil {
		return nil, err
	}
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return &users, nil

}

func login(a authProto.AuthenticationService, username, password string) (*authProto.LoginResponse, error) {
	loginRequest := &authProto.LoginRequest{username, password}
	resp, err := a.Login(context.Background(), loginRequest)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func register(a authProto.AuthenticationService, user *userProto.User) (*authProto.RegisterResponse, error) {
	resp, err := a.Register(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
