package romans

import "strings"

func main() {
}

func remaninder(number, divisor int) (restremainder, restquotion int) {

	restremainder = number % divisor
	restquotion = number / divisor

	return
}
func Convert(number int) (romanresult string) {
	//var romanresult string

	var romanEquivalences = []struct {
		arabic int    // input
		roman  string // expected result
	}{
		{arabic: 5, roman: "V"},
		{arabic: 4, roman: "IV"},
		{arabic: 1, roman: "I"},
	}

	for _, numbersmap := range romanEquivalences {
		numberofsymbols, decimalremainder := remaninder(number, numbersmap.arabic)
		number = decimalremainder
		romanresult += strings.Repeat(numbersmap.roman, numberofsymbols)
	}

	return
}
