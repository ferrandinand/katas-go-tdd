package main

import (
	"fmt"
	"strconv"
)

func main() {
	fill(100)
}

func fill(n int) []int {
	maxnumber := make([]int, n)

	for i := range maxnumber {
		maxnumber[i] = i
	}
	return maxnumber
}

func divisibleby3(n int) (bool, string) {

	if n%3 == 0 {
		return true, "fizz"
	}
	return false, strconv.Itoa(n)
}

func divisibleby5(n int) (bool, string) {

	if n%5 == 0 {
		return true, "buzz"
	}
	return false, strconv.Itoa(n)

}

func divisibleby3and5(n int) (bool, string) {
	isdivisibleby3, _ := divisibleby3(n)
	isdivisibleby5, _ := divisibleby5(n)

	if isdivisibleby3 && isdivisibleby5 {
		return true, "fizzbuzz"
	}
	return false, strconv.Itoa(n)
}

func transformFizzOrBuzz(arrayofnumbers []int) []string {
	var transformedArray []string

	for i := 0; i < len(arrayofnumbers); i++ {
		fmt.Println(arrayofnumbers[i])
	}

	// 	for i := 0; i < len(arrayofnumbers); i++ {

	// 		transformedArray[i] = strconv.Itoa(arrayofnumbers[i])

	// 	}

	return transformedArray
}
