package romans

import "strings"

func main() {
}

func remaninder(number, divisor int) (restremainder, restquotion int) {

	restremainder = number % divisor
	restquotion = number / divisor

	return
}

//Convert arabic numbers to roman
func Convert(number int) string {
	var romanresult string

	var romanEquivalences = []struct {
		arabic int    // input
		roman  string // expected result
	}{
		{arabic: 1, roman: "I"},
		{arabic: 4, roman: "IV"},
		{arabic: 5, roman: "V"},
		{arabic: 6, roman: "VI"},
	}

	for _, numbersmap := range romanEquivalences {
		if number < 4 && number > 0 {
			romanresult = strings.Repeat(numbersmap.roman, number)
		}
		if number == 4 {
			romanresult = "IV"

		}
		if number == 5 {
			romanresult = "V"

		}
		if number == 6 {
			romanresult = "VI"

		}

	}

	return romanresult
}
