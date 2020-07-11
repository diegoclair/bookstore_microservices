package main

import (
	"net/http"
	"time"

	"github.com/diegoclair/microservice_items/server"
	"github.com/diegoclair/microservice_items/service"
)

func main() {

	svc := service.New( /*db*/ )
	srv := server.InitServer(svc)

	server := &http.Server{
		Addr:         "localhost:3002",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      srv,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
