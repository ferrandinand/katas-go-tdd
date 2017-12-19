package doom

import (
	"math/rand"
)

var player string

type attackEvent struct {
	daemonId int
}
type daemon struct {
	id         int
	numAttacks int
}

type gameState struct {
	daemons []daemon
}

func (state gameState) getDaemon(id int) *daemon {
	for index, daemon := range state.daemons {
		if daemon.id == id {
			return &state.daemons[index]
		}
	}
	return nil
}

func attack() attackEvent {
	return attackEvent{daemonId: rand.Intn(2000)}
}

func processEvent(state gameState, event attackEvent) gameState {
	selectedDaemon := state.getDaemon(event.daemonId)
	selectedDaemon.numAttacks = 1
	return state
}
