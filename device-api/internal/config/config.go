package config

import (
	"os"

	"github.com/rs/zerolog"
)

type DeviceAPIConfig struct {
	Logger *zerolog.Logger
	Port   int `yaml:"DEVICE_API_DEFAULT_PORT"`
}

func NewDeviceAPIConfig(logger *zerolog.Logger) *DeviceAPIConfig {
	c := &DeviceAPIConfig{
		Logger: logger,
	}
	return c
}

var configPath = `../templates/configs/device-api.yaml`

func (c *DeviceAPIConfig) FromEnv() *DeviceAPIConfig {
	values, err := LoadTemplateConfigMap(configPath) // load the yaml file
	if err != nil {
		c.Logger.Err(err).Msg("failed to load config map with err")
		os.Exit(1)
	}
	c.Port = values.Port
	return c
}
