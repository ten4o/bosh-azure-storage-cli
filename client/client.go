package client

import (
	"os"
)

type AzBlobstore struct {
	storageClient StorageClient
}

func New(storageClient StorageClient) (AzBlobstore, error) {
	return AzBlobstore{storageClient: storageClient}, nil
}

func (client *AzBlobstore) Put(sourceFile *os.File, destPath string) error {

	_, err := client.storageClient.Upload(sourceFile, destPath)

	return err
}
