package client

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"time"
)

type StorageResponse struct {
	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
	ClientRequestID *string

	// ContentMD5 contains the information returned from the Content-MD5 header response.
	ContentMD5 []byte

	// Date contains the information returned from the Date header response.
	Date *time.Time

	// ETag contains the information returned from the ETag header response.
	ETag *azcore.ETag

	// EncryptionKeySHA256 contains the information returned from the x-ms-encryption-key-sha256 header response.
	EncryptionKeySHA256 *string

	// EncryptionScope contains the information returned from the x-ms-encryption-scope header response.
	EncryptionScope *string

	// IsServerEncrypted contains the information returned from the x-ms-request-server-encrypted header response.
	IsServerEncrypted *bool

	// LastModified contains the information returned from the Last-Modified header response.
	LastModified *time.Time

	// RequestID contains the information returned from the x-ms-request-id header response.
	RequestID *string

	// Version contains the information returned from the x-ms-version header response.
	Version *string

	// VersionID contains the information returned from the x-ms-version-id header response.
	VersionID *string
}
