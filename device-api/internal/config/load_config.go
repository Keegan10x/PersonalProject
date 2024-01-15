package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func LoadTemplateConfigMap(path string) (*DeviceAPIConfig, error) {

	// read the YAML file
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// unmarshal the YAML content into a Config struct
	var config DeviceAPIConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
