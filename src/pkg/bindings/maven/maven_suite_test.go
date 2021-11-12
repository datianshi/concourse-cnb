package maven_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMaven(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Maven Suite")
}
