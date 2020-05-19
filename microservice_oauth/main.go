package main

import (
	"log"

	"github.com/diegoclair/microservice_oauth/data"
	"github.com/diegoclair/microservice_oauth/server"
	"github.com/diegoclair/microservice_oauth/service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	log.Println("Reading the intial configs...")

	db, err := data.Connect()
	if err != nil {
		panic(err)
	}

	svc := service.New(db)
	server := server.InitServer(svc)

	if err := server.Run(":3001"); err != nil {
		panic(err)
	}
}
