package doom

import (
	"fmt"
	"math/rand"
)

var player string

func attack() string {
	demon := []string{"daemon 1", "daemon 2", "daemon 3"}

	ran := rand.Intn(2)
	fmt.Print(rand.Intn(2000))
	return demon[ran]
}
