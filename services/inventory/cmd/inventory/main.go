package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	inventoryService "github.com/yash91989201/ecomm-monorepo/services/inventory"
)

type Config struct {
	DatabaseUrl string `envconfig:"DATABASE_URL" required:"true"`
	ServiceUrl  string `envconfig:"SERVICE_URL" required:"true"`
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	var r inventoryService.Repository
	retry.ForeverSleep(5*time.Second, func(_ int) (err error) {
		r, err = inventoryService.NewMysqlRepository(cfg.DatabaseUrl)

		if err != nil {
			log.Println(err)
		}

		return
	})

	defer r.Close()

	log.Printf("Inventory service started at %s", cfg.ServiceUrl)

	s := inventoryService.New(r)
	log.Fatal(inventoryService.Start(s, cfg.ServiceUrl))
}
