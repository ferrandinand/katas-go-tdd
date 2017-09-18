package romans

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRomans(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Romans Suite")

}

var _ = Describe("Roman numbers", func() {
	Context("Test of the first 6 numbers", func() {
		It("convert 1 to roman", func() {
			Expect(Convert(1)).Should(BeEquivalentTo("I"))
		})
		It("convert 2 to roman", func() {
			Expect(Convert(2)).Should(BeEquivalentTo("II"))
		})
		It("convert 3 to roman", func() {
			Expect(Convert(3)).Should(BeEquivalentTo("III"))
		})
		It("convert 4 to roman", func() {
			Expect(Convert(4)).Should(BeEquivalentTo("IV"))
		})
		It("convert 5 to roman", func() {
			Expect(Convert(5)).Should(BeEquivalentTo("V"))
		})
		It("convert 6 to roman", func() {
			Expect(Convert(6)).Should(BeEquivalentTo("VI"))
		})
	})
})
