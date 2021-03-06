package main

import (
	"github.com/CESARBR/knot-babeltower/internal/config"
	"github.com/CESARBR/knot-babeltower/pkg/controllers"
	"github.com/CESARBR/knot-babeltower/pkg/interactors"
	"github.com/CESARBR/knot-babeltower/pkg/network"
	"github.com/CESARBR/knot-babeltower/pkg/server"

	"github.com/CESARBR/knot-babeltower/pkg/logging"
)

func main() {
	config := config.Load()
	logrus := logging.NewLogrus(config.Logger.Level)

	logger := logrus.Get("Main")
	logger.Info("Starting KNoT Babeltower")

	userProxy := network.NewUserProxy(logrus.Get("UserProxy"), config.Users.Hostname, config.Users.Port)
	createUser := interactors.NewCreateUser(logrus.Get("CreateUser"), userProxy)
	userController := controllers.NewUserController(logrus.Get("Controller"), createUser)
	server := server.NewServer(config.Server.Port, logrus.Get("Server"), userController)
	server.Start()
}
