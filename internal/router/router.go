package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/abdul-ghaffar01/api-gateway/internal/config"
)

type Router struct {
	Routes map[string][]string
}

func New(config config.Config) (*Router, error) {
	// Creating an empty routing table
	var router Router = Router{Routes: make(map[string][]string)}

	for _, v := range config.Routes {
		key := v.Path

		_, exists := router.Routes[key]
		if exists {
			// Check if there are any duplicate routes
			alreadyRegisteredRoutes := router.Routes[key]
			newRoutesToBeRegistered := v.Methods
			for _, method := range newRoutesToBeRegistered {
				for _, methodRegistered := range alreadyRegisteredRoutes {
					if strings.EqualFold(method, methodRegistered) {
						return nil, fmt.Errorf("Dulicate registration for path \"%v\"", v.Path)
					}
				}
			}
			// Already a method exists, need to append these new
			router.Routes[key] = append(router.Routes[key], v.Methods...)
		} else {
			router.Routes[key] = v.Methods
		}
	}

	return &router, nil
}



func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 1. Find path

    // 2. Verify method

    // 3. Rate limit

    // 4. Authenticate

    // 5. Proxy request

    // 6. Return response
}