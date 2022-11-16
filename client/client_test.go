package client_test

import (
	"github.com/mvach/bosh-azure-storage-cli/client"
	"github.com/mvach/bosh-azure-storage-cli/client/clientfakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Client", func() {

	It("put file uploads to a blob", func() {
		storageClient := clientfakes.FakeStorageClient{}

		azBlobstore, err := client.New(&storageClient)
		Expect(err).ToNot(HaveOccurred())

		file, _ := os.CreateTemp("", "tmpfile")

		azBlobstore.Put(file, "target/blob")

		Expect(storageClient.UploadCallCount()).To(Equal(1))
		source, dest := storageClient.UploadArgsForCall(0)

		Expect(source).To(Equal(file))
		Expect(dest).To(Equal("target/blob"))
	})

	It("get blob downloads to a file", func() {
		storageClient := clientfakes.FakeStorageClient{}

		azBlobstore, err := client.New(&storageClient)
		Expect(err).ToNot(HaveOccurred())

		file, _ := os.CreateTemp("", "tmpfile")

		azBlobstore.Get("source/blob", file)

		Expect(storageClient.DownloadCallCount()).To(Equal(1))
		source, dest := storageClient.DownloadArgsForCall(0)

		Expect(source).To(Equal("source/blob"))
		Expect(dest).To(Equal(file))
	})

})
