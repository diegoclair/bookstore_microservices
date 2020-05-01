package main

import (
	"net/http"

	"github.com/diegoclair/microservice_items/server"
	"github.com/diegoclair/microservice_items/service"
)

func main() {

	svc := service.New( /*db*/ )
	srv := server.InitServer(svc)

	server := &http.Server{
		Handler: srv,
		Addr:    "127.0.0.1:3003",
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
