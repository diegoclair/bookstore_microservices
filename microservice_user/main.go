package main

import (
	"github.com/diegoclair/bookstore_users-api/data"
	"github.com/diegoclair/bookstore_users-api/logger"
	"github.com/diegoclair/bookstore_users-api/server"
	"github.com/gin-gonic/gin"

	"github.com/diegoclair/bookstore_users-api/service"
)

func main() {
	logger.Info("Reading the intial configs...")

	db, err := data.Connect()
	if err != nil {
		panic(err)
	}
	svc := service.New(db)
	server := initializeServer(svc)

	server.Run(":3000")
}

func initializeServer(svc *service.Service) *gin.Engine {

	srv := server.InitServer(svc)
	logger.Info("About to start the application...")

	return srv
}
