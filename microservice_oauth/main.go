package main

import (
	"github.com/diegoclair/microservice_oauth/data"
	"github.com/diegoclair/microservice_oauth/server"
	"github.com/diegoclair/microservice_oauth/service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	db, err := data.Connect()
	if err != nil {
		panic(err)
	}

	svc := service.New(db)
	server := server.InitServer(svc)

	server.Run(":3001")
}
