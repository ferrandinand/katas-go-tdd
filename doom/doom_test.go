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
			attackEvent := attackEvent{daemonId: 35}
			gameState := gameState{daemonId: 35, numAttacks: 0}

			newGameState := processEvent(gameState, attackEvent)

			Expect(newGameState.daemonId).To(Equal(35))
			Expect(newGameState.numAttacks).To(Equal(1))
		})

		/*It("EventAttack generates a new game state", func() {
			attackEvent := attackEvent{daemonId: 35}
			attackEvent := attackEvent{daemonId: 37}

			gameState := gameState{daemons:[{}]

			newGameState := processEvent(gameState, attackEvent)

			Expect(newGameState.daemonId).To(Equal(35))
			Expect(newGameState.numAttacks).To(Equal(1))
		})*/

	})
})
