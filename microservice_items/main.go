package main

import (
	"net/http"
	"time"

	"github.com/diegoclair/go_utils-lib/logger"
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

	logger.Info("About to start the application on port 3002...")

	if err := server.ListenAndServe(); err != nil {
		logger.Error("Error to start", err)
		panic(err)
	}

}
