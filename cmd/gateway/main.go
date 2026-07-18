package main

import (
	"github.com/abdul-ghaffar01/api-gateway/internal/config"
	"flag"
	"fmt"
	"os"
)


func main(){
	path := "config.yaml"

	if env := os.Getenv("CONFIG_PATH"); env != "" {
		path = env
	}

	configPath := flag.String("config", path, "Configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)

	fmt.Println(*configPath)

	if err != nil {
		panic(err)
	}

	// Create server using cfg
	fmt.Print(cfg)
}