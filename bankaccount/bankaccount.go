package bankaccount

import (
	"strconv"
	"time"
)

var balanceTotal float64
var input_movement float64 = 300.0

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
	return time.Now().Format("d-m-Y")
}

func outputMovement() string {
	movement := strconv.FormatFloat(input_movement, 'E', -1, 64)
	balanceStatus := strconv.FormatFloat(balance(), 'E', -1, 64)

	myBalance := date() + movement + balanceStatus
	return myBalance
}
