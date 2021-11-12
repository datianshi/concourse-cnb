package maven_test

import (
	"io/ioutil"

	"github.com/datianshi/concourse-cnb/config/pkg/bindings/maven"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var result string = `
<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0 https://maven.apache.org/xsd/settings-1.0.0.xsd">
  <servers>
    <server>
      <id>buildpack_maven_repo</id>
      <username>test-account</username>
      <password>test-password</password>
      <configuration></configuration>
    </server>
  </servers>
</settings>
`

var _ = Describe("Settings", func() {

	var (
		username string
		password string
		repoId   string
		settings maven.Settings
	)
	BeforeEach(func() {
		username = "test-account"
		password = "test-password"
		repoId = "buildpack_maven_repo"
		settings = *maven.NewSettings(username, password, repoId)
	})

	It("Should generate correct setting content", func() {
		c, err := settings.GenerateContent()
		Ω(err).Should(BeNil())
		b, err := ioutil.ReadAll(c)
		Ω(err).Should(BeNil())
		Ω(result).Should(Equal(string(b)))
	})
})
