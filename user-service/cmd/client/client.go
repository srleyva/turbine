package main

import (
	"fmt"
	micro "github.com/micro/go-micro"
	proto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
	"io"
)

func main() {
	service := micro.NewService(micro.Name("user.client"))
	service.Init()
	user := proto.NewUserService("user-service", service.Client())
	_, err := user.CreateUser(context.TODO(), &proto.User{Username: "sleyva", Password: "Test"})
	if err != nil {
		fmt.Printf("err running rpc command: %v", err)
	}

	users, err := user.GetUsers(context.TODO(), &proto.UsersRequest{All: false, Count: 1})
	if err != nil {
		fmt.Println("error: %v", err)
	}
	for {
		user, err := users.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("%v.ListFeatures(_) = _, %v", user, err)
		}
		fmt.Println(user)
	}
}
