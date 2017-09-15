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
	Context("Test of the first 3 numbers", func() {
		It("1 to roman", func() {
			Expect(Convert(1)).Should(BeEquivalentTo("I"))
		})
		It("2 to roman", func() {
			Expect(Convert(2)).Should(BeEquivalentTo("II"))
		})
		It("3 to roman", func() {
			Expect(Convert(3)).Should(BeEquivalentTo("III"))
		})
	})
})
