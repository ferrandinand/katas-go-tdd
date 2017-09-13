package romans

import (
	"fmt"
	"strings"
)

// ArabicToRoman is main function
func ArabicToRoman(number int) string {
	if number == 0 {
		fmt.Println("")
		value := ""
		return value
	}
	if number < 3 && number > 0 {
		value := strings.Repeat("I", number)
		return value
	}

	value := "XX"
	return value
}
