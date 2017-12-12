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
		It("Attacks a random demon", func() {
			Expect(attack()).To(Equal("daemon 1"))
		})

	})
})
