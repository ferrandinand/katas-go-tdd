package doom

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDoom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Doom")
}

var _ = Describe("Doom", func() {

	Context("", func() {
		It("Check daemonID is not nil", func() {
			Expect(attack().daemonId).ShouldNot(BeNil())
		})

	})
	Context("", func() {
		It("EventAttack generates a new game state", func() {
			attack := attackEvent{daemonId: 35}

			gameState := gameState{daemons: []daemon{daemon{id: 35, numAttacks: 0}}}

			finalGameState := processEvent(gameState, attack)

			Expect(finalGameState.daemons[0].id).To(Equal(35))
			Expect(finalGameState.daemons[0].numAttacks).To(Equal(1))
		})

		It("EventAttack generates a new game state", func() {
			firstAttack := attackEvent{daemonId: 35}
			secondAttack := attackEvent{daemonId: 37}

			gameState := gameState{daemons: []daemon{daemon{id: 35, numAttacks: 0}, daemon{id: 37, numAttacks: 0}}}

			intermediateGameState := processEvent(gameState, firstAttack)
			finalGameState := processEvent(intermediateGameState, secondAttack)

			Expect(finalGameState.daemons[0].id).To(Equal(35))
			Expect(finalGameState.daemons[0].numAttacks).To(Equal(1))

			Expect(finalGameState.daemons[1].id).To(Equal(37))
			Expect(finalGameState.daemons[1].numAttacks).To(Equal(1))

		})

	})
})
