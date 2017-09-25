package romans

import "strings"

func main() {
}

func getremainderandquotient(number, divisor int) (remainder, quotient int) {

	remainder = number % divisor
	quotient = number / divisor

	return
}

//Convertarabictoroman is a func to arabic numbers to roman numbers
func Convertarabictoroman(number int) (romanresult string) {
	var romanequivalences = []struct {
		arabic int
		roman  string
	}{
		{arabic: 50, roman: "L"},
		{arabic: 40, roman: "XL"},
		{arabic: 10, roman: "X"},
		{arabic: 9, roman: "IX"},
		{arabic: 5, roman: "V"},
		{arabic: 4, roman: "IV"},
		{arabic: 1, roman: "I"},
	}

	for _, arabicromanmap := range romanequivalences {
		decimalremainder, numberofsymbols := getremainderandquotient(number, arabicromanmap.arabic)
		number = decimalremainder
		romanresult += strings.Repeat(arabicromanmap.roman, numberofsymbols)
	}

	return
}
