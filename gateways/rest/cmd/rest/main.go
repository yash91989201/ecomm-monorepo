package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/yash91989201/ecomm-monorepo/gateways/rest"
)

type Config struct {
	Port                int    `envconfig:"PORT" required:"true"`
	InventoryServiceUrl string `envconfig:"INVENTORY_SERVICE_URL" required:"true"`
}

func main() {

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	server, err := rest.NewServer(ctx, cfg.InventoryServiceUrl, cfg.Port)
	if err != nil {
		fmt.Println("Failed to create server")
	}

	go func() {
		if err := server.Start(); err != nil {
			fmt.Println("Failed to start rest api gateway")
		}
	}()

	<-ctx.Done()

	log.Println("Closing service client connections.")
	server.CloseServiceClients()

	log.Println("Shutting down the server...")

	timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(timeout); err != nil {
		log.Printf("Error during shutdown: %v", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}
