package bankaccount

import (
	"strconv"
	"time"
)

var balanceTotal float64
var input_movement float64 = 500.00

func deposit(received float64) {
	balanceTotal += received
}

func resetBalance() {
	balanceTotal = 0
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
	movement := formatAmount(input_movement)
	balanceStatus := formatAmount(balance())
	myBalance := "date || credit || balance\n" + date() + " || " + movement + " || " + balanceStatus
	return myBalance
}
