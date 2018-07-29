package login

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	authProto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	userProto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"net/http"
)

const (
	// TODO Store in Vault
	Key = "secret"
)

type Handler struct {
	Auth authProto.AuthenticationService
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

	log.Info("User: ", u)

	response, err := register(h.Auth, u)
	if err != nil {
		if err.Error() == "empty username or password" {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err}
		}
		return err
	}

	return c.JSON(http.StatusCreated, response)

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
