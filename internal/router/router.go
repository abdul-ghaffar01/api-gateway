package router

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/abdul-ghaffar01/api-gateway/internal/config"
)

type Router struct {
	Routes map[string]map[string]config.Route
}

func New(cfg config.Config) (*Router, error) {
	router := Router{
		Routes: make(map[string]map[string]config.Route),
	}

	for _, route := range cfg.Routes {
		path := route.Path

		// Create a method map for this path if it doesn't exist.
		if _, exists := router.Routes[path]; !exists {
			router.Routes[path] = make(map[string]config.Route)
		}

		// Register each method.
		for _, method := range route.Methods {
			method = strings.ToUpper(method)

			if _, exists := router.Routes[path][method]; exists {
				return nil, fmt.Errorf(
					"duplicate registration for path %q and method %q",
					path,
					method,
				)
			}

			router.Routes[path][method] = route
		}
	}

	return &router, nil
}

// handles all the requests
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// For demo
	fmt.Println(req)

	client := &http.Client{}
	methods, ok := r.Routes[req.URL.Path]
	if !ok {
		http.Error(w, "Route not found", http.StatusNotFound)
		return
	}

	route, ok := methods[req.Method]
	if !ok {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	upstreamReq, err := http.NewRequest(
		req.Method,
		route.Upstream,
		req.Body,
	)

	if err != nil {
		http.Error(w, "Upstream unavailable", http.StatusBadGateway)
		return
	}

upstreamReq.Header = req.Header.Clone()
upstreamReq.Header.Del("Accept-Encoding")
	resp, err := client.Do(upstreamReq)

	if err != nil {
		http.Error(w, "Upstream unavailable", http.StatusBadGateway)
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

	// Tasks:
	// 1. Find path

	// 2. Verify method

	// 3. Rate limit

	// 4. Authenticate

	// 5. Proxy request

	// 6. Return response
}
