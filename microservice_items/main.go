package main

import (
	"log"
	"net/http"
	"time"

	"github.com/diegoclair/microservice_items/data"
	"github.com/diegoclair/microservice_items/server"
	"github.com/diegoclair/microservice_items/service"
)

func main() {

	db, err := data.Connect()
	if err != nil {
		panic(err)
	}

	svc := service.New(db)
	srv := server.InitServer(svc)

	server := &http.Server{
		Addr:         ":3002",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      srv,
	}

	log.Println("Listening on port 3002...")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
