package env_test

import (
	"fmt"
	"os"

	"github.com/datianshi/concourse-cnb/config/pkg/env"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Env", func() {
	var (
		envDir string = fmt.Sprintf("%s%s", os.TempDir(), "env")
		prefix string = "BUILD_ENV_"
	)

	Describe("create environment files", func() {
		BeforeEach(func() {
			os.Setenv("BUILD_ENV_BP_TEST_VAR1", "VALUE1")
			os.Setenv("BUILD_ENV_BP_TEST_VAR2", "VALUE2")
		})
		It("Should create correct key and file", func() {
			err := env.EnvBuildpack(prefix, envDir)
			Ω(err).Should(BeNil())

			filePath1 := fmt.Sprintf("%s/%s", envDir, "BP_TEST_VAR1")
			_, err = os.Stat(filePath1)
			Ω(err).Should(BeNil())
			content, _ := os.ReadFile(filePath1)
			Ω("VALUE1").Should(Equal(string(content)))

			filePath2 := fmt.Sprintf("%s/%s", envDir, "BP_TEST_VAR2")
			_, err = os.Stat(filePath2)
			Ω(err).Should(BeNil())
			content, _ = os.ReadFile(filePath2)
			Ω("VALUE2").Should(Equal(string(content)))
		})
		AfterEach(func() {
			os.Unsetenv("BUILD_ENV_BP_TEST_VAR1")
			os.Unsetenv("BUILD_ENV_BP_TEST_VAR1")
		})
	})
})
