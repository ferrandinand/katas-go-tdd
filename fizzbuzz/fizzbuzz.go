package main

import "fmt"

func main() {
	fill(100)
}

func fill(n int) []int {
	maxnumber := make([]int, n)

	for i := range maxnumber {
		maxnumber[i] = i
	}
	fmt.Println(maxnumber)
	return maxnumber
}

func divisibleby3(n int) bool {

	if n%3 == 0 {
		return true
	}
	return false
}

func divisibleby5(n int) bool {

	if n%5 == 0 {
		return true
	}
	return false

}

func divisibleby3and5(n int) bool {

	if divisibleby3(n) && divisibleby5(n) {
		return true
	}
	return false

}

// func transformFizzOrBuzz(arrayofnumbers []int) []string {
// 	var transformedArray []string
// 	for i, value := range arrayofnumbers {
// 		fmt.Println(value)
// 		transformedArray[i] = strconv.Itoa(value)

// 	}

// 	for i := 0; i < len(arrayofnumbers); i++ {

// 		transformedArray[i] = strconv.Itoa(arrayofnumbers[i])

// 	}

// 	return transformedArray
// }
