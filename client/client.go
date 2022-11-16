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

func (client *AzBlobstore) Put(source *os.File, dest string) error {

	_, err := client.storageClient.Upload(source, dest)

	return err
}

func (client *AzBlobstore) Get(source string, dest *os.File) error {

	_, err := client.storageClient.Download(source, dest)

	return err
}
