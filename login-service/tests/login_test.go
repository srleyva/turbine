package main

import (
	"github.com/labstack/echo"
	"github.com/micro/go-micro/client"
	authProto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	"github.com/srleyva/turbine/login-service/pkg/login"
	userProto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Mocks
type mockAuthService struct {
}

func (m *mockAuthService) Login(ctx context.Context, req *authProto.LoginRequest, opts ...client.CallOption) (*authProto.LoginResponse, error) {
	resp := &authProto.LoginResponse{Token: "token", Username: req.Username}
	return resp, nil
}

func (m *mockAuthService) Register(ctx context.Context, req *userProto.User, opts ...client.CallOption) (*authProto.RegisterResponse, error) {
	resp := &authProto.RegisterResponse{true, req}
	return resp, nil
}

// Set Up
var (
	mockauth     = mockAuthService{}
	loginHandler = login.Handler{&mockauth}
	userJSON     = `{"username":"sleyva","password":"test"}`
)

func TestLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := loginHandler.Login(c); err != nil {
		t.Errorf("Error returned where not expected: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Wrong Status Returned: Wanted %d, Got: %d", http.StatusCreated, rec.Code)
	}

	if rec.Body.String() != `{"token":"token","username":"sleyva"}` {
		t.Errorf("Response Malformed: %s", rec.Body.String())
	}
}

func TestRegister(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := loginHandler.Register(c); err != nil {
		t.Errorf("Error returned where not expected: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Wrong Status Returned: Wanted %d, Got: %d", http.StatusCreated, rec.Code)
	}

	if rec.Body.String() != `{"created":true,"user":{"username":"sleyva","password":"test"}}` {
		t.Errorf("Response Malformed: %s", rec.Body.String())
	}
}
