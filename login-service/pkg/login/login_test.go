package login_test

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	authMocks "github.com/srleyva/turbine/authentication-service/mocks"
	authProto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	. "github.com/srleyva/turbine/login-service/pkg/login"
	userMocks "github.com/srleyva/turbine/user-service/mocks"
	userProto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	userJSON     = `{"username":"thisdude","password":"test"}`
	UsersRequest = `{"all":true}`
)

type tester struct {
	MockAuthService *authMocks.MockAuthenticationService
	MockUserService *userMocks.MockUserService
	Ctrl            *gomock.Controller
	Handler         *Handler
}

func newTester(t *testing.T) *tester {
	ctrl := gomock.NewController(t)
	mockAuth := authMocks.NewMockAuthenticationService(ctrl)
	mockUser := userMocks.NewMockUserService(ctrl)
	handler := &Handler{mockAuth, mockUser}
	loginTester := &tester{mockAuth, mockUser, ctrl, handler}
	return loginTester
}

func TestLogin(t *testing.T) {

	tester := newTester(t)
	handler := tester.Handler

	tester.MockAuthService.EXPECT().Login(
		context.Background(),
		&authProto.LoginRequest{Username: "thisdude", Password: "test"}).Return(
		&authProto.LoginResponse{
			Token:    "token",
			Username: "sleyva"}, nil)

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := handler.Login(c); err != nil {
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

	tester := newTester(t)
	handler := tester.Handler
	tester.MockAuthService.EXPECT().Register(
		context.Background(),
		&userProto.User{Username: "thisdude", Password: "test"}).Return(
		&authProto.RegisterResponse{
			Created: true,
			User: &userProto.User{
				UID:      "this-dude-is-cool",
				Username: "thisdude",
				Password: ""}}, nil)

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := handler.Register(c); err != nil {
		t.Errorf("Error returned where not expected: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Wrong Status Returned: Wanted %d, Got: %d", http.StatusCreated, rec.Code)
	}

	if rec.Body.String() != `{"created":true,"user":{"username":"thisdude","UID":"this-dude-is-cool"}}` {
		t.Errorf("Response Malformed: %s", rec.Body.String())
	}
}

func TestGetAllUsers(t *testing.T) {
	tester := newTester(t)
	handler := tester.Handler
	mockStream := userMocks.NewMockUserService_GetUsersService(tester.Ctrl)

	tester.MockUserService.EXPECT().GetUsers(
		context.Background(),
		&userProto.UsersRequest{All: true}).Return(mockStream, nil).Return(mockStream, nil)

	firstRequest := mockStream.EXPECT().Recv().Return(&userProto.User{UID: "this-dude-is-cool", Username: "thisdude"}, nil)
	mockStream.EXPECT().Recv().After(firstRequest).Return(nil, io.EOF)

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(UsersRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := handler.GetAllUsers(c); err != nil {
		t.Errorf("Error returned where not expected: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Wrong Status Returned: Wanted %d, Got: %d", http.StatusCreated, rec.Code)
	}

	if rec.Body.String() != `[{"username":"thisdude","UID":"this-dude-is-cool"}]` {
		t.Errorf("Response Malformed: %s", rec.Body.String())
	}
}
