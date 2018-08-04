package main

import (
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	authProto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	userProto "github.com/srleyva/turbine/user-service/proto/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/srleyva/turbine/login-service/pkg/login"
	"net/http"
)

var (
	// service to call
	authService string
	userService string
	appDir      string
)

func main() {
	// Initialize microservice clients
	service := micro.NewService()
	service.Init(micro.Flags(cli.StringFlag{
		Name:        "auth_service",
		Value:       "go.micro.srv.authentication",
		Destination: &authService,
	},
		cli.StringFlag{
			Name:        "user_service",
			Value:       "go.micro.srv.user",
			Destination: &userService,
		},
		cli.StringFlag{
			Name:        "app_dir",
			Value:       "/app/wwwroot",
			Destination: &appDir,
		}))
	auth := authProto.NewAuthenticationService(authService, service.Client())
	user := userProto.NewUserService(userService, service.Client())

	// Initialize Server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(login.Key),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/auth/login" || c.Path() == "/auth/register" || c.Path() == "/static/*" || c.Path() == "/" {
				return true
			}
			return false
		},
	}))

	fs := http.FileServer(http.Dir(appDir))
	e.GET("/", echo.WrapHandler(fs))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))

	h := login.Handler{auth, user}
	e.GET("/auth/users", h.GetAllUsers)
	e.POST("/auth/login", h.Login)
	e.POST("/auth/register", h.Register)
	e.Logger.Fatal(e.Start(":1323"))
}
