package main

import (
	"fmt"
	micro "github.com/micro/go-micro"
	proto "github.com/srleyva/turbine/user-service/proto/user"
	context "golang.org/x/net/context"
)

func main() {
	service := micro.NewService(micro.Name("user.client"))
	service.Init()
	user := proto.NewUserService("user-service", service.Client())
	rsp, err := user.CreateUser(context.TODO(), &proto.User{Username: "sleyva", Password: "Test"})
	if err != nil {
		fmt.Printf("err running rpc command: %v", err)
	}

	fmt.Printf("User: %s", rsp.User.Username)
}
