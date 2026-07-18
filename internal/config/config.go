package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// To validate the configurations
func validateConfig(config Config) error {
	// Checking if port is valid
	if config.Server.Port <= 0 || config.Server.Port > 65535 {
		return fmt.Errorf("Invalid port number passed")
	}

	// Checking if routes are defined
	if len(config.Routes) == 0 {
		return fmt.Errorf("Routes are not defined in configuration file")
	}

	 // Validating each route
	for _, route := range config.Routes {

		// Checking if the upstream or path are passed
		if route.Upstream == ""{
			return fmt.Errorf("upstream is not defined in route \"%v\"", route.Name)
		}

		// Checking if the upstream or path are passed
		if route.Path == "" {
			return fmt.Errorf("path is not defined in route \"%v\"", route.Name)
		}

		// Checking for http methods
		if len(route.Methods) == 0 {
			return fmt.Errorf("method is not defined for route \"%v\"", route.Name)
		}

		// if all the passed methods are valid
		for _, method := range route.Methods {
			var VALID_METHODS = [10]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "TRACE", "CONNECT", "QUERY"}
			isValid := false

			for _, v := range VALID_METHODS {
				if strings.ToUpper(method) == v {
					isValid = true
					break
				}
			}

			if !isValid {
				return fmt.Errorf("Invalid method name \"%v\" in route \"%v\"", method, route.Name)
			}
		}
	}

	return nil
}


// To load the config file
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

	// checking if the configuration is correct
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return &config, nil
}
