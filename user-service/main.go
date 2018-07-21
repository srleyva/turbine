package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/srleyva/turbine/user-service/handler"
	"github.com/srleyva/turbine/user-service/subscriber"

	example "github.com/srleyva/turbine/user-service/proto/example"
)

var version string

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user-service"),
		micro.Version(version),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.user-service", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.user-service", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
