package main

import (
	"main/location-calculator/internal"
	"main/location-calculator/internal/config"
	"main/location-calculator/internal/server"
	"main/services/logger"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	// define dependencies
	router := mux.NewRouter()
	logger := logger.NewLogger()
	c := config.NewLocationCalculatorConfig(logger).FromEnv()

	// register routes
	internal.PublicRoutes(router, nil)

	// start server
	svr := server.NewServer(router, logger, c)
	err := svr.Run()
	if err != nil {
		logger.Error().Err(err).Int("tried port", c.Port) // log out error
		os.Exit(2)
	}
}
