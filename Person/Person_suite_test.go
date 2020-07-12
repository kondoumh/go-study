package Person_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPerson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Person Suite")
}
