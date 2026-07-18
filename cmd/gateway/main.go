package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abdul-ghaffar01/api-gateway/internal/config"
	"github.com/abdul-ghaffar01/api-gateway/internal/router"
	"github.com/abdul-ghaffar01/api-gateway/internal/server"
)


func main(){
	path := "config.yaml"

	if env := os.Getenv("CONFIG_PATH"); env != "" {
		path = env
	}

	configPath := flag.String("config", path, "Configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	// config load/validation error
	if err != nil {
		panic(err)
	}

	// Initialize routing table
	router, errRouter := router.New(*cfg)
	if errRouter != nil {
		panic(errRouter)
	}

	fmt.Println(*router)

	// Create server using cfg, and router
	server := server.New(*cfg, router)

	// Starting the server to listen for requests
	errServer := server.Start()

	if errServer != nil {
		panic(errServer)
	}

	fmt.Print(server)
}