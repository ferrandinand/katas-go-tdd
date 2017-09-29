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
		Array []int
		//transformedArray []string
	)

	BeforeEach(func() {
		Array = fill(100)
		//transformedArray = transformFizzOrBuzz(Array)
	})

	Context("Create list of 100", func() {
		It("Create an array of 100 items", func() {
			Expect(len(Array)).Should(BeEquivalentTo(100))
		})
	})

	Context("Numbers divisible by 3 show Fizz", func() {
		It("3 must return Fizz", func() {
			Expect(divisibleby3(3)).To(Equal(true))
		})
	})

	Context("Numbers divisible by 5 show Buzz", func() {
		It("5 must return Fizz", func() {
			Expect(divisibleby5(5)).To(Equal(true))
		})
	})

	Context("Numbers divisible by 3 and 5 show FizzBuzz", func() {
		It("15 must return FizzBuzz", func() {
			Expect(divisibleby3and5(15)).To(Equal(true))
		})
	})
	//Context("Check expected values", func() {
	//	It("Contains Fizz Buzz FizzBuzz or integer", func() {
	//		Expect((transformedArray)).Should()
	//	})
	//})
})
