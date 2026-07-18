package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)


// To validate the configurations
func validateConfig(config Config) error {
	// Checking if port is valid
	if config.Server.Port < 0 || config.Server.Port > 65535 {
		return fmt.Errorf("Invalid port number passed")
	}

	// Checking if routes are defined
	if len(config.Routes) == 0 {
		return fmt.Errorf("Routes are not defined in configuration file")
	}

	 // Validating each route
	for _, route := range config.Routes {

		// Checking if the base_url or path are passed
		if route.Base_url == ""{
			return fmt.Errorf("base_url is not defined in route %v", route.Name)
		}

		// Checking if the base_url or path are passed
		if route.Path == "" {
			return fmt.Errorf("path is not defined in route %v", route.Name)
		}
	}

	return nil
}

func Load(path string) (*Config, error) {
	// Reading the config file
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	// Config
	var config Config

	// Parse the data into structs
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return &config, nil
}
