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
	)

	BeforeEach(func() {
		Array := fill(100)
	})

	Context("Create list of 100", func() {
		It("Create an array of 100 items", func() {
			Expect(len(Array)).Should(BeEquivalentTo(100))
		})
	})

})

// type name struct {
// 	X string
// }

// func fill(n int) []int {
// 	var maxnumber []int

// 	for i := 1; i <= n; i++ {
// 		maxnumber[i] = i
// 		fmt.Println("THis is %v", maxnumber[i])
// 	}
// 	return maxnumber
// }

// func TestFizzBuzz(t *testing.T) {
// 	cases := []struct {
// 		in, want string
// 	}{
// 		{"1", "1"},
// 		{"2", "2"},
// 		{"3", "fizz"},
// 		{"4", "4"},
// 		{"5", "buzz"},
// 		{"6", "fizz"},
// 		{"7", "7"},
// 		{"8", "8"},
// 		{"9", "fizz"},
// 		{"10", "buzz"},
// 		{"11", "11"},
// 		{"12", "fizz"},
// 		{"13", "13"},
// 		{"14", "14"},
// 		{"15", "buzz"},
// 		{"18", "fizz"},
// 		{"20", "buzz"},
// 		{"25", "buzz"},
// 	}

// 	}
//}

// testing example https://golang.org/doc/code.html#Testing
