package romans

import "strings"

func main() {

}

//Convert arabic numbers to roman
func Convert(number int) string {
	var romanresult string
	// if number == 1 {
	// 	romanresult = "I"
	// } else {
	// 	romanresult = "II"
	// }

	if number < 4 && number > 0 {
		romanresult = strings.Repeat("I", number)
	}
	return romanresult
}
