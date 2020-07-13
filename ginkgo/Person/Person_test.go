package Person_test

import (
	"errors"

	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Person", func() {
	Context("Test for Name", func() {
		It("Test for SetName", func() {
			p := &Person{}
			p.SetName("Alice")
		})
		It("Test for GetName (Normal)", func() {
			p := &Person{}
			p.SetName("Bob")
			name, err := p.GetName()
			Expect(err).To(BeNil())
			Expect(name).To(Equal("Bob"))
		})
		It("Test for GetName (Error)", func() {
			p := &Person{}
			_, err := p.GetName()
			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(errors.New("Name is not set")))
		})
	})
})
