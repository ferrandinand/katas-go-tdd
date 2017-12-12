package bankaccount

import (
	"strconv"
	"time"
)

var balanceTotal float64
var input_movement float64 = 500.00 // convert to function

var banksettelment = []struct {
	date    string
	credit  float64
	balance float64
}{
	{date: "01/04/2014", credit: 1000.00, balance: 1000.00},
	{date: "02/04/2014", credit: 500.00, balance: 1500.00},
}

func deposit(received float64) {
	balanceTotal += received
}

func balance() float64 {
	return balanceTotal
}

func date() string {
	return time.Now().Format("2/01/2006")
}

func formatAmount(amount float64) string {
	return strconv.FormatFloat(amount, 'f', 2, 64)
}

func outputMovement() string {
	myBalance := "date || credit || balance\n" + banksettelment[0].date + " || " + strconv.FormatFloat(banksettelment[0].credit, 'f', 2, 64) + " || " + strconv.FormatFloat(banksettelment[0].balance, 'f', 2, 64)
	return myBalance
}
