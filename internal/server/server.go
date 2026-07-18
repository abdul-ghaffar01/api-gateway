package server

import (
	"fmt"
	"net/http"

	"github.com/abdul-ghaffar01/api-gateway/internal/config"
	"github.com/abdul-ghaffar01/api-gateway/internal/router"
)

type Server struct {
	httpServer *http.Server
}

// Creating the new server
func New(config config.Config, r *router.Router) (*Server, error) {
	// Create a new server
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", config.Server.Port),
		Handler: r,
	}

	server := Server{
		httpServer: httpServer,
	}
	return &server, nil
}