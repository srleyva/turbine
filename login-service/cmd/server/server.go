package main

import (
	"fmt"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	authProto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	userProto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
)

var (
	// service to call
	serviceName string
)

func main() {
	service := micro.NewService()

	service.Init(micro.Flags(cli.StringFlag{
		Name:        "service_name",
		Value:       "go.micro.srv.authentication",
		Destination: &serviceName,
	}))

	auth := authProto.NewAuthenticationService(serviceName, service.Client())

	if err := register(auth, userProto.User{Username: "sleyva", Password: "test"}); err != nil {
		fmt.Println(err)
	}

	if err := login(auth, "sleyva", "test4"); err != nil {
		fmt.Printf("err: %v", err)
		return
	}
}

func login(a authProto.AuthenticationService, username, password string) error {
	loginRequest := &authProto.LoginRequest{username, password}
	resp, err := a.Login(context.Background(), loginRequest)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return err
	}
	fmt.Println(resp)
	return nil
}

func register(a authProto.AuthenticationService, user userProto.User) error {
	resp, err := a.Register(context.Background(), &user)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return err
	}
	fmt.Println(resp)
	return nil
}
