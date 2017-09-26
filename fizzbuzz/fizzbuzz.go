package main

import "fmt"
import "strconv"

func fizzBuzz(max int) string {

	results := ""

	for i := 1; i <= max; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	return results
}

func main() {
	fill(100)
}

func fill(n int) []int {
	maxnumber := make([]int, n)

	for i := 0; i < n; i++ {
		maxnumber[i] = i + 1
	}
	return maxnumber
}

func divisibleby3(n int) string {

	if n%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)

}

func transformFizzOrBuzz(arrayofnumbers []int) []string {
	var transformedArray []string
	for i, value := range arrayofnumbers {
		fmt.Println(value)
		transformedArray[i] = strconv.Itoa(value)

	}

	// for i := 0; i < len(arrayofnumbers); i++ {

	// 	transformedArray[i] = strconv.Itoa(arrayofnumbers[i])

	// }

	return transformedArray
}
