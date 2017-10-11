package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFizzBuzz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FizzBuzz")
}

var _ = Describe("FizzBuzz", func() {
	var (
		arrayofnumbers   []int
		transformedArray []string
	)
	BeforeEach(func() {
		arrayofnumbers = fill(100)
		transformedArray = transformFizzOrBuzz(arrayofnumbers)
	})

	AfterEach(func() {
		arrayofnumbers = nil
	})

	Context("Create list of 100", func() {
		It("Create an array of 100 items", func() {
			Expect(len(arrayofnumbers)).Should(BeEquivalentTo(100))
		})
	})

	Context("Numbers divisible by 3 show Fizz", func() {

		It("3 must return Fizz", func() {
			divisible, number := divisibleby3(3)
			Expect(divisible).To(Equal(true))
			Expect(number).To(Equal("fizz"))
		})
		It("4 must return Fizz", func() {
			divisible, number := divisibleby3(4)
			Expect(divisible).To(Equal(false))
			Expect(number).To(Equal("4"))
		})
	})

	Context("Numbers divisible by 5 show Buzz", func() {
		It("5 must return buzz", func() {
			divisible, number := divisibleby5(5)
			Expect(divisible).To(Equal(true))
			Expect(number).To(Equal("buzz"))
		})
		It("7 must return buzz", func() {
			divisible, number := divisibleby5(7)
			Expect(divisible).To(Equal(false))
			Expect(number).To(Equal("7"))
		})
	})

	Context("Numbers divisible by 3 and 5 show FizzBuzz", func() {
		It("15 must return FizzBuzz", func() {
			divisible, number := divisibleby3and5(15)
			Expect(divisible).To(Equal(true))
			Expect(number).To(Equal("fizzbuzz"))
		})
		It("16 must return 16", func() {
			divisible, number := divisibleby3and5(16)
			Expect(divisible).To(Equal(false))
			Expect(number).To(Equal("16"))
		})
	})

	Context("Check expected values", func() {
		It("Contains Fizz if it divisible by 3", func() {
			Expect(transformedArray[2]).To(BeEquivalentTo("Fizz"))
		})
	})
})
