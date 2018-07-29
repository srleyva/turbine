package main

import (
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	authProto "github.com/srleyva/turbine/authentication-service/proto/authentication"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/srleyva/turbine/login-service/pkg/login"
)

var (
	// service to call
	serviceName string
)

func main() {
	// Initialize microservice clients
	service := micro.NewService()
	service.Init(micro.Flags(cli.StringFlag{
		Name:        "service_name",
		Value:       "go.micro.srv.authentication",
		Destination: &serviceName,
	}))
	auth := authProto.NewAuthenticationService(serviceName, service.Client())

	// Initialize Server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(login.Key),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/login" || c.Path() == "/register" || c.Path() == "/" {
				return true
			}
			return false
		},
	}))

	h := login.Handler{auth}
	e.Static("/", "dist")
	e.POST("/login", h.Login)
	e.POST("/register", h.Register)
	e.Logger.Fatal(e.Start(":1323"))
}
