package bindings_test

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/datianshi/concourse-cnb/config/pkg/bindings"
	"github.com/datianshi/concourse-cnb/config/pkg/bindings/bindingsfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bindings", func() {

	var (
		content    *bindingsfakes.FakeContent = &bindingsfakes.FakeContent{}
		bindingDir string                     = fmt.Sprintf("%s%s", os.TempDir(), "bindings")
		binding    bindings.Binding
	)

	Context("with content", func() {
		BeforeEach(func() {
			binding = *bindings.NewBinding("my-binding", "maven", bindingDir, content)
			content.GenerateContentStub = func() (io.Reader, error) {
				b := new(bytes.Buffer)
				b.WriteString("This is binding content")
				return b, nil
			}
			content.NameStub = func() string {
				return "settings.xml"
			}
		})
		It("Should generate correct file structure", func() {

			err := binding.CreateBinding()
			Ω(err).Should(BeNil())

			//Binding name dir should be created
			bindingNameDir := fmt.Sprintf("%s/%s", bindingDir, "my-binding")
			_, err = os.Stat(bindingNameDir)
			Ω(err).Should(BeNil())

			//Should have a file type
			bindingTypeFile := fmt.Sprintf("%s/%s", bindingNameDir, "type")
			_, err = os.Stat(bindingTypeFile)
			Ω(err).Should(BeNil())
			c, _ := os.ReadFile(bindingTypeFile)
			Ω("maven").Should(Equal(string(c)))

			//The content should be generated
			bindingContentFile := fmt.Sprintf("%s/%s", bindingNameDir, "settings.xml")
			_, err = os.Stat(bindingContentFile)
			Ω(err).Should(BeNil())
			c, _ = os.ReadFile(bindingContentFile)
			Ω("This is binding content").Should(Equal(string(c)))
		})
	})

	Context("no content", func() {
		BeforeEach(func() {
			binding = *bindings.NewBinding("my-binding", "maven", bindingDir, &bindings.EmptyContent{})
		})

		It("Should not throw error", func() {
			err := binding.CreateBinding()
			Ω(err).Should(BeNil())
		})
	})
})
