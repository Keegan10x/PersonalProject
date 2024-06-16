package server

import (
	"fmt"
	"main/location-calculator/internal/config"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type server struct {
	router *mux.Router
	logger *zerolog.Logger
	config *config.LocationCalculatorConfig
}

func NewServer(router *mux.Router, logger *zerolog.Logger, config *config.LocationCalculatorConfig) *server {
	s := &server{
		router: router,
		logger: logger,
		config: config,
	}
	return s
}

func (s *server) Run() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	err := http.ListenAndServe(addr, s.router) // listening message
	if err != nil {
		return err
	}
	s.logger.Info().Msg("location-calculator started") // message
	return nil
}
