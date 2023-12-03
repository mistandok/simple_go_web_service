package main

import (
	"net/http"
	clients "simple_web_app/internal/clients"
	"simple_web_app/internal/loggers"
	services "simple_web_app/internal/services"
	storages "simple_web_app/internal/storages"
)

func main() {
	logger := services.LoggerAdapter(loggers.LogOutput)
	dataStore := storages.NewSimpleDataStore()
	logic := services.NewSimpleLogic(logger, dataStore)
	controller := clients.NewController(logger, logic)
	http.HandleFunc("/hello", controller.SayHello)
	http.ListenAndServe(":8080", nil)
}
