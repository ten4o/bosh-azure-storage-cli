package client

import (
	"fmt"
	"github.com/mvach/bosh-azure-storage-cli/blob"
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

	_, err := client.storageClient.Upload(source, dest)

	return err
}

func (client *AzBlobstore) Get(source string, dest *os.File) error {

	_, err := client.storageClient.Download(source, dest)

	return err
}

func (client *AzBlobstore) Delete(dest string) error {

	_, err := client.storageClient.Delete(dest)

	return err
}

func (client *AzBlobstore) Exists(dest string) (blob.ExistenceState, error) {

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
