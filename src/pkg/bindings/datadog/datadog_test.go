package datadog_test

import (
	"io/ioutil"

	"github.com/datianshi/concourse-cnb/config/pkg/bindings/datadog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var result string = `
api_key: my_api_key
metadata_collectors:
  - name: 'resources'
    interval: 60
log_file: AGENT_LOG_FILE
apm_config:
  enabled: true
  log_file: TRACE_LOG_FILE
`

var _ = Describe("Datadog", func() {
	var (
		apiKey string
		dg     datadog.DataDog
	)
	BeforeEach(func() {
		apiKey = "my_api_key"
		dg = *datadog.NewDataDog(apiKey)
	})

	It("Should generate correct setting content", func() {
		c, err := dg.GenerateContent()
		Ω(err).Should(BeNil())
		b, err := ioutil.ReadAll(c)
		Ω(err).Should(BeNil())
		Ω(result).Should(Equal(string(b)))
	})
})
