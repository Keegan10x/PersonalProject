package config

import (
	"os"

	"github.com/rs/zerolog"
)

var configPath = `../templates/configs/device-api.yaml`

type LocationCalculatorConfig struct {
	Logger *zerolog.Logger
	Port   int `yaml:"LOCATION_CALCULATOR_DEFAULT_PORT"`
}

func NewLocationCalculatorConfig(logger *zerolog.Logger) *LocationCalculatorConfig {
	c := &LocationCalculatorConfig{
		Logger: logger,
	}
	return c
}

func (c *LocationCalculatorConfig) FromEnv() *LocationCalculatorConfig {
	values, err := LoadTemplateConfigMap(configPath) // load the yaml file
	if err != nil {
		c.Logger.Err(err).Msg("failed to load config map with err")
		os.Exit(1)
	}
	c.Port = values.Port
	return c
}
