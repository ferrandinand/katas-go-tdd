package roman
//Fix should be package roman_test

import (
    "testing"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

func TestRoman(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Roman Numbers Suite")
}

var _ = Describe("Roman numbers", func() {
    Context("initially", func() {
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
