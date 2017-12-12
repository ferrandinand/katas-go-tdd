package doom

import (
	"math/rand"
)

var player string

type attackEvent struct {
	daemonId int
}

type gameState struct {
	daemonId   int
	numAttacks int
}

func attack() attackEvent {
	return attackEvent{daemonId: rand.Intn(2000)}
}

func processEvent(state gameState, event attackEvent) gameState {
	state.numAttacks = 1
	return state
}
