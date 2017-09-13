package romans

import (
	"bytes"
	"fmt"
)

//romans gets a number and returns a roman equivalent
func romans(number int) string {

	RomanMap := map[int]string{
		1: "I",
		5: "V",
	}

	var RomanArray []string
	var DecimalArray []int

	for k, j := range RomanMap {
		DecimalArray = append(DecimalArray, k)
		RomanArray = append(RomanArray, j)
	}

	LastNumber := DecimalArray[len(DecimalArray)-1:][0]
	var buffer bytes.Buffer

	//decimalIndex := len(DecimalArray) - 1
	for number > 3 {

		if number < LastNumber {
			return "IV"

		}
		return "V"
	}

	buffer.WriteString(lastThree(number))
	fmt.Println(buffer.String())
	return buffer.String()
}

//lastThree inserts last III
func lastThree(number int) string {
	var buffer bytes.Buffer

	for i := number; i > 0; i-- {
		buffer.WriteString("I")
	}

	fmt.Println(buffer.String())
	return buffer.String()
}
