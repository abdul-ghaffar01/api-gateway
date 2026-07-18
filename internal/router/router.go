package router

import (

	"github.com/abdul-ghaffar01/api-gateway/internal/config"
)

type Router struct {
	Routes map[string][]string
}

func New(config config.Config) (*Router, error) {
	var router Router

	for _, v := range config.Routes{
		key := v.Base_url + v.Path

		_, exists := router.Routes[key]
		if exists {
			// TODO: Check if there are any duplicate route
			// Already a method exists, need to append these new
			router.Routes[key] = append(router.Routes[key], v.Methods...)
		}
		router.Routes[key] = v.Methods
	}

	return &router, nil
}