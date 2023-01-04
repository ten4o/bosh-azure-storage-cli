package client

import (
	"context"
	"fmt"
	azBlob "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/mvach/bosh-azure-storage-cli/blob"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/mvach/bosh-azure-storage-cli/config"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . StorageClient
type StorageClient interface {
	Upload(
		source io.ReadSeekCloser,
		dest string,
	) (StorageUploadResponse, error)

	Download(
		source string,
		dest *os.File,
	) (int64, error)

	Delete(
		dest string,
	) (StorageDeleteResponse, error)

	Exists(
		dest string,
	) (blob.ExistenceState, error)

	SignedUrl(
		dest string,
		expiration time.Duration,
	) (string, error)
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
	source io.ReadSeekCloser,
	dest string,
) (StorageUploadResponse, error) {

	blobURL := fmt.Sprintf("%s/%s", dsc.serviceURL, dest)

	log.Println(fmt.Sprintf("Uploading %s", blobURL))
	client, err := blockblob.NewClientWithSharedKeyCredential(blobURL, dsc.credential, nil)
	if err != nil {
		return StorageUploadResponse{}, err
	}

	resp, err := client.Upload(context.Background(), source, nil)

	return StorageUploadResponse{
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

func (dsc DefaultStorageClient) Download(
	source string,
	dest *os.File,
) (int64, error) {

	blobURL := fmt.Sprintf("%s/%s", dsc.serviceURL, source)

	log.Println(fmt.Sprintf("Downloading %s", blobURL))
	client, err := blockblob.NewClientWithSharedKeyCredential(blobURL, dsc.credential, nil)
	if err != nil {
		return -1, err
	}

	resp, err := client.DownloadFile(context.Background(), dest, nil)

	return resp, err
}

func (dsc DefaultStorageClient) Delete(
	dest string,
) (StorageDeleteResponse, error) {

	blobURL := fmt.Sprintf("%s/%s", dsc.serviceURL, dest)

	log.Println(fmt.Sprintf("Deleting %s", blobURL))
	client, err := blockblob.NewClientWithSharedKeyCredential(blobURL, dsc.credential, nil)
	if err != nil {
		return StorageDeleteResponse{}, err
	}

	resp, err := client.Delete(context.Background(), nil)

	return StorageDeleteResponse{
		ClientRequestID: resp.ClientRequestID,
		Date:            resp.Date,
		RequestID:       resp.RequestID,
		Version:         resp.Version,
	}, err
}

func (dsc DefaultStorageClient) Exists(
	dest string,
) (blob.ExistenceState, error) {

	blobURL := fmt.Sprintf("%s/%s", dsc.serviceURL, dest)

	log.Println(fmt.Sprintf("Checking if blob: %s exists", blobURL))
	client, err := blockblob.NewClientWithSharedKeyCredential(blobURL, dsc.credential, nil)
	if err != nil {
		return blob.ExistenceUnknown, err
	}

	_, err = client.BlobClient().GetProperties(context.Background(), nil)
	if strings.Contains(err.Error(), "RESPONSE 404") {
		return blob.NotExisting, nil
	}
	if err != nil {
		return blob.ExistenceUnknown, err
	}

	return blob.Existing, nil
}

func (dsc DefaultStorageClient) SignedUrl(
	dest string,
	expiration time.Duration,
) (string, error) {

	blobURL := fmt.Sprintf("%s/%s", dsc.serviceURL, dest)

	log.Println(fmt.Sprintf("Getting signed url for blob %s", blobURL))
	client, err := azBlob.NewClientWithSharedKeyCredential(blobURL, dsc.credential, nil)
	if err != nil {
		return "", err
	}

	url, err := client.GetSASURL(sas.BlobPermissions{Read: true, Create: true}, time.Now(), time.Now().Add(expiration))
	if err != nil {
		return "", err
	}

	return url, err
}
