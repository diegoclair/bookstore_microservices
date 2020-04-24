package main

import (
	"github.com/diegoclair/bookstore_oauth-api/data"
	"github.com/diegoclair/bookstore_oauth-api/server"
	"github.com/diegoclair/bookstore_oauth-api/service"
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
	server := initializeServer(svc)

	server.Run(":3001")
}

func initializeServer(svc *service.Service) *gin.Engine {

	srv := server.InitServer(svc)

	return srv
}
