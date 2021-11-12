package bindings_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBindings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bindings Suite")
}
