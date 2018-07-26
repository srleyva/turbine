package main

import (
	log "github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	auth "github.com/srleyva/turbine/authentication-service/pkg/authentication"
	proto "github.com/srleyva/turbine/authentication-service/proto/authentication"
	"github.com/srleyva/turbine/user-service/pkg/logging"
	userProto "github.com/srleyva/turbine/user-service/proto/user"
	micrologrus "github.com/tudurom/micro-logrus"
)

func main() {
	golog := logrus.New()
	golog.Formatter = logging.Formatter
	logger := micrologrus.NewMicroLogrus(golog)
	log.SetLogger(logger)
	authService := micro.NewService(
		micro.Name("authentication-service"),
		micro.WrapHandler(logging.LogWrapper),
	)

	userClient := userProto.NewUserService("user-service", authService.Client())

	authService.Init()
	proto.RegisterAuthenticationServiceHandler(authService.Server(), &auth.Service{userClient})

	if err := authService.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
