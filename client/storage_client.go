package client

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/mvach/bosh-azure-storage-cli/config"
	"io"
	"log"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . StorageClient
type StorageClient interface {
	Upload(
		sourceFile io.ReadSeekCloser,
		destPath string,
	) (StorageResponse, error)
}

type DefaultStorageClient struct {
	credential *azblob.SharedKeyCredential
	serviceURL string
}

func NewStorageClient(storageConfig config.AZStorageConfig) (StorageClient, error) {
	credential, err := azblob.NewSharedKeyCredential(storageConfig.AccountName, storageConfig.AccountKey)
	if err != nil {
		return nil, err
	}

	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", storageConfig.AccountName, storageConfig.ContainerName)

	return DefaultStorageClient{credential: credential, serviceURL: serviceURL}, nil
}

func (dsc DefaultStorageClient) Upload(
	sourceFile io.ReadSeekCloser,
	destPath string) (StorageResponse, error) {

	blobURL := fmt.Sprintf("%s/%s", dsc.serviceURL, destPath)

	log.Println(fmt.Sprintf("Uploading %s", blobURL))
	client, err := blockblob.NewClientWithSharedKeyCredential(blobURL, dsc.credential, nil)
	if err != nil {
		return StorageResponse{}, err
	}

	resp, err := client.Upload(context.Background(), sourceFile, nil)

	return StorageResponse{
		ClientRequestID:     resp.ClientRequestID,
		ContentMD5:          resp.ContentMD5,
		Date:                resp.Date,
		ETag:                resp.ETag,
		EncryptionKeySHA256: resp.EncryptionKeySHA256,
		EncryptionScope:     resp.EncryptionScope,
		IsServerEncrypted:   resp.IsServerEncrypted,
		LastModified:        resp.LastModified,
		RequestID:           resp.RequestID,
		Version:             resp.Version,
		VersionID:           resp.VersionID,
	}, err
}
