package client

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type AzBlobstore struct {
	storageClient StorageClient
}

func New(storageClient StorageClient) (AzBlobstore, error) {
	return AzBlobstore{storageClient: storageClient}, nil
}

func (client *AzBlobstore) Put(source *os.File, dest string) error {

	return client.storageClient.Upload(source, dest)
}

func (client *AzBlobstore) Get(source string, dest *os.File) error {

	return client.storageClient.Download(source, dest)
}

func (client *AzBlobstore) Delete(dest string) error {

	return client.storageClient.Delete(dest)
}

func (client *AzBlobstore) Exists(dest string) (bool, error) {

	return client.storageClient.Exists(dest)
}

func (client *AzBlobstore) Sign(dest string, action string, expiration time.Duration) (string, error) {
	action = strings.ToUpper(action)
	switch action {
	case "GET", "PUT":
		return client.storageClient.SignedUrl(dest, expiration)
	default:
		return "", fmt.Errorf("action not implemented: %s", action)
	}
}
