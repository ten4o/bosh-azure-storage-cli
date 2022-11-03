package main

import (
	"github.com/mvach/bosh-azure-storage-cli/client"
	"github.com/mvach/bosh-azure-storage-cli/config"
	"log"
)

func main() {
	azConfig, err := config.NewFromReader(nil)
	if err != nil {
		log.Fatalln(err)
	}

	azClient, err := client.NewAZClient(azConfig)
	if err != nil {
		log.Fatalln(err)
	}
}
