package config_test

import (
	"bytes"
	"errors"
	"github.com/mvach/bosh-azure-storage-cli/config"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	It("sets host and port", func() {
		configJson := []byte(`{"host": "foo-host", "port": "1234"}`)
		configReader := bytes.NewReader(configJson)

		config, err := config.NewFromReader(configReader)

		Expect(err).ToNot(HaveOccurred())
		Expect(config.Host).To(Equal("foo-host"))
		Expect(config.Port).To(Equal("1234"))
	})

	Context("when the configuration file cannot be read", func() {
		It("returns an error", func() {
			f := explodingReader{}

			_, err := config.NewFromReader(f)
			Expect(err).To(MatchError("explosion"))
		})
	})

})

type explodingReader struct{}

func (e explodingReader) Read([]byte) (int, error) {
	return 0, errors.New("explosion")
}
