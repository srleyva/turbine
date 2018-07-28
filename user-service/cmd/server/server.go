package main

import (
	log "github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"github.com/srleyva/turbine/user-service/pkg/logging"
	user "github.com/srleyva/turbine/user-service/pkg/user"
	proto "github.com/srleyva/turbine/user-service/proto/user"
	micrologrus "github.com/tudurom/micro-logrus"
)

func main() {
	golog := logrus.New()
	golog.Formatter = logging.Formatter
	logger := micrologrus.NewMicroLogrus(golog)
	log.SetLogger(logger)
	datastore := &user.DataStore{}
	userService := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.WrapHandler(logging.LogWrapper),
	)

	userService.Init()
	proto.RegisterUserServiceHandler(userService.Server(), &user.Service{datastore})
	if err := userService.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
