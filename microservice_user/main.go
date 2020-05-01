package main

import (
	"github.com/diegoclair/microservice_user/data"
	"github.com/diegoclair/microservice_user/logger"
	"github.com/diegoclair/microservice_user/server"

	"github.com/diegoclair/microservice_user/service"
)

func main() {
	logger.Info("Reading the intial configs...")

	db, err := data.Connect()
	if err != nil {
		panic(err)
	}
	svc := service.New(db)
	server := server.InitServer(svc)
	logger.Info("About to start the application...")

	server.Run(":3000")
}
