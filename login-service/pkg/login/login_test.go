package login_test

import (
	"errors"
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

// Set Up
var (
	userJSON = `{"username":"sleyva","password":"test"}`
)

// Helpers
type loginTester struct {
	Handler   *Handler
	UserMocks *userMocks.MockUserService
	AuthMocks *authMocks.MockAuthenticationService
	Ctrl      *gomock.Controller
}

func newLoginTester(t *testing.T) *loginTester {
	ctrl := gomock.NewController(t)
	userMock := userMocks.NewMockUserService(ctrl)
	authMock := authMocks.NewMockAuthenticationService(ctrl)
	handler := &Handler{authMock, userMock}
	tester := &loginTester{handler, userMock, authMock, ctrl}
	return tester
}

// Tests
func TestLogin(t *testing.T) {
	tester := newLoginTester(t)
	defer tester.Ctrl.Finish()

	e := echo.New()

	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("test user is returned correctly when no error", func(t *testing.T) {
		tester.AuthMocks.EXPECT().Login(
			context.Background(),
			&authProto.LoginRequest{Username: "sleyva", Password: "test"}).Return(
			&authProto.LoginResponse{Token: "this", Username: "sleyva"}, nil)

		if err := tester.Handler.Login(c); err != nil {
			t.Errorf("Error returned where not expected: %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("Wrong Status Returned: Wanted %d, Got: %d", http.StatusCreated, rec.Code)
		}

		if rec.Body.String() != `{"token":"this","username":"sleyva"}` {
			t.Errorf("Response Malformed: %s", rec.Body.String())
		}
	})
}

func TestRegister(t *testing.T) {
	tester := newLoginTester(t)
	defer tester.Ctrl.Finish()

	tester.AuthMocks.EXPECT().Register(
		context.Background(),
		&userProto.User{Username: "sleyva", Password: "test"}).Return(
		&authProto.RegisterResponse{
			Created: true,
			User:    &userProto.User{Username: "sleyva", Password: "", UID: "this-is-an-id"}}, nil)

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := tester.Handler.Register(c); err != nil {
		t.Errorf("Error returned where not expected: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Wrong Status Returned: Wanted %d, Got: %d", http.StatusCreated, rec.Code)
	}

	if rec.Body.String() != `{"created":true,"user":{"username":"sleyva","UID":"this-is-an-id"}}` {
		t.Errorf("Response Malformed: %s", rec.Body.String())
	}
}

func TestGetAllUsers(t *testing.T) {
	tester := newLoginTester(t)
	defer tester.Ctrl.Finish()

	mockStream := userMocks.NewMockUserService_GetUsersService(tester.Ctrl)

	tester.UserMocks.EXPECT().GetUsers(context.Background(), &userProto.UsersRequest{All: true}).Return(mockStream, nil)
	firstCall := mockStream.EXPECT().Recv().Return(&userProto.User{Username: "Sleyva", Password: "test"}, nil)
	mockStream.EXPECT().Recv().After(firstCall).Return(nil, io.EOF)

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/auth/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := tester.Handler.GetAllUsers(c); err != nil {
		t.Errorf("Error returned where not expected: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Wrong Status Returned: Wanted %d, Got: %d", http.StatusCreated, rec.Code)
	}

}
