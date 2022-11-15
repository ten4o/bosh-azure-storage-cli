package main

import (
	"flag"
	"log"
	"os"

	"github.com/mvach/bosh-azure-storage-cli/client"
	"github.com/mvach/bosh-azure-storage-cli/config"
)

func main() {

	configPath := flag.String("c", "", "configuration path")
	flag.Parse()

	configFile, err := os.Open(*configPath)
	if err != nil {
		log.Fatalln(err)
	}

	azConfig, err := config.NewFromReader(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	storageClient, err := client.NewStorageClient(azConfig)
	if err != nil {
		log.Fatalln(err)
	}

	azBlobstore, err := client.New(storageClient)
	if err != nil {
		log.Fatalln(err)
	}

	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) < 2 {
		log.Fatalf("Expected at least two arguments got %d\n", len(nonFlagArgs))
	}

	cmd := nonFlagArgs[0]

	switch cmd {
	case "put":
		if len(nonFlagArgs) != 3 {
			log.Fatalf("Put method expected 3 arguments got %d\n", len(nonFlagArgs))
		}
		src, dst := nonFlagArgs[1], nonFlagArgs[2]

		var sourceFile *os.File
		sourceFile, err = os.Open(src)
		if err != nil {
			log.Fatalln(err)
		}

		defer sourceFile.Close()

		err = azBlobstore.Put(sourceFile, dst)

	default:
		log.Fatalf("unknown command: '%s'\n", cmd)
	}
}
