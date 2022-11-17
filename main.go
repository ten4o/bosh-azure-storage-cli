package main

import (
	"flag"
	"github.com/mvach/bosh-azure-storage-cli/blob"
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

	blobstoreClient, err := client.New(storageClient)
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

		err = blobstoreClient.Put(sourceFile, dst)
		if err != nil {
			log.Fatalln(err)
		}

	case "get":
		if len(nonFlagArgs) != 3 {
			log.Fatalf("Get method expected 3 arguments got %d\n", len(nonFlagArgs))
		}
		src, dst := nonFlagArgs[1], nonFlagArgs[2]

		var dstFile *os.File
		dstFile, err = os.Create(dst)
		if err != nil {
			log.Fatalln(err)
		}

		defer dstFile.Close()

		err = blobstoreClient.Get(src, dstFile)
		if err != nil {
			log.Fatalln(err)
		}

	case "delete":
		if len(nonFlagArgs) != 2 {
			log.Fatalf("Delete method expected 2 arguments got %d\n", len(nonFlagArgs))
		}

		err = blobstoreClient.Delete(nonFlagArgs[1])
		if err != nil {
			log.Fatalln(err)
		}

	case "exists":
		if len(nonFlagArgs) != 2 {
			log.Fatalf("Existing method expected 2 arguments got %d\n", len(nonFlagArgs))
		}

		existsState, err := blobstoreClient.Exists(nonFlagArgs[1])
		if err != nil {
			log.Fatalln(err)
		}

		// If the object exists the exit status is 0, otherwise it is 3
		// We are using `3` since `1` and `2` have special meanings
		if existsState == blob.NotExisting {
			os.Exit(3)
		}

	default:
		log.Fatalf("unknown command: '%s'\n", cmd)
	}
}
