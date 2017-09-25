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
		It("Converts the arabic number 1 to roman", func() {
			Expect(Convertarabictoroman(1)).Should(BeEquivalentTo("I"))
		})
		It("Converts the arabic number 2 to roman", func() {
			Expect(Convertarabictoroman(2)).Should(BeEquivalentTo("II"))
		})
		It("Converts the arabic number 3 to roman", func() {
			Expect(Convertarabictoroman(3)).Should(BeEquivalentTo("III"))
		})
		It("Converts the arabic number 4 to roman", func() {
			Expect(Convertarabictoroman(4)).Should(BeEquivalentTo("IV"))
		})
		It("Converts the arabic number 5 to roman", func() {
			Expect(Convertarabictoroman(5)).Should(BeEquivalentTo("V"))
		})
		It("Converts the arabic number 6 to roman", func() {
			Expect(Convertarabictoroman(6)).Should(BeEquivalentTo("VI"))
		})
		It("Converts the arabic number 9 to roman", func() {
			Expect(Convertarabictoroman(9)).Should(BeEquivalentTo("IX"))
		})
		It("Converts the arabic number 12 to roman", func() {
			Expect(Convertarabictoroman(12)).Should(BeEquivalentTo("XII"))
		})
		It("Converts the arabic number 15 to roman", func() {
			Expect(Convertarabictoroman(15)).Should(BeEquivalentTo("XV"))
		})
		It("Converts the arabic number 19 to roman", func() {
			Expect(Convertarabictoroman(19)).Should(BeEquivalentTo("XIX"))
		})
		It("Converts the arabic number 40 to roman", func() {
			Expect(Convertarabictoroman(40)).Should(BeEquivalentTo("XL"))
		})
		It("Converts the arabic number 41 to roman", func() {
			Expect(Convertarabictoroman(41)).Should(BeEquivalentTo("XLI"))
		})
		It("Converts the arabic number 50 to roman", func() {
			Expect(Convertarabictoroman(50)).Should(BeEquivalentTo("L"))
		})
	})
})
