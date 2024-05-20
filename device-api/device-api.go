package main

import (
	"main/device-api/internal/config"
	"main/device-api/internal/handlers"
	"main/device-api/internal/server"
	"main/services/logger"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// define dependencies
	router := mux.NewRouter()
	logger := logger.NewLogger()
	c := config.NewDeviceAPIConfig(logger).FromEnv()

	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	// start server
	svr := server.NewServer(router, logger, c)
	err := svr.Run()
	if err != nil {
		logger.Error().Err(err).Int("tried port", c.Port) // log out error
		os.Exit(2)
	}
}
