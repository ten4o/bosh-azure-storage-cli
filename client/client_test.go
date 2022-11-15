package client_test

import (
	"github.com/mvach/bosh-azure-storage-cli/client"
	"github.com/mvach/bosh-azure-storage-cli/client/clientfakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Config", func() {

	It("put", func() {
		storageClient := clientfakes.FakeStorageClient{}

		azBlobstore, err := client.New(&storageClient)
		Expect(err).ToNot(HaveOccurred())

		file, _ := os.CreateTemp("", "tmpfile")

		azBlobstore.Put(file, "target/file.txt")

		Expect(storageClient.UploadCallCount()).To(Equal(1))
		sourceFile, destPath := storageClient.UploadArgsForCall(0)

		Expect(sourceFile).To(Equal(file))
		Expect(destPath).To(Equal("target/file.txt"))
	})

})
