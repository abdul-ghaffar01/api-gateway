package main

import (
	"flag"
	"os"
	"github.com/abdul-ghaffar01/api-gateway/internal/config"
)


func main(){
	path := "config.yaml"

	if env := os.Getenv("CONFIG_PATH"); env != "" {
		path = env
	}

	configPath := flag.String("config", path, "Configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
}