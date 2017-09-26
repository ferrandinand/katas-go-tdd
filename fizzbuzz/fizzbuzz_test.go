package main

import "testing"
import . "github.com/onsi/ginkgo"
import . "github.com/onsi/gomega"

func TestFizzBuzz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FizzBuzz")
}

var _ = Describe("FizzBuzz", func() {
	var (
		Array            []int
		transformedArray []string
	)

	BeforeEach(func() {
		Array = fill(100)
		transformedArray = transformFizzOrBuzz(Array)
	})

	Context("Create list of 100", func() {
		It("Create an array of 100 items", func() {
			Expect(len(Array)).Should(BeEquivalentTo(100))
		})
	})
	Context("Numbers divisible by 3 show Fizz", func() {
		It("3 must return Fizz", func() {
			Expect(divisibleby3(3)).To(Equal("Fizz"))
		})
	})
	Context("Check expected values", func() {
		It("3 must return Fizz", func() {
			Expect(len(transformedArray)).Should(BeEquivalentTo(100))
		})
	})
})
