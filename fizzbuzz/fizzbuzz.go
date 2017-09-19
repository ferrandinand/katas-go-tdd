package main

import "fmt"

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
	//	fizzBuzz(5)
	fill(100)
}

func fill(n int) []int {
	maxnumber := make([]int, n)

	for i := 0; i < n; i++ {
		maxnumber[i] = i + 1
		fmt.Println(maxnumber[i])
	}
	return maxnumber
}
