package romans

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRomans(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Romans Suite")

}

var _ = Describe("Roman numbers", func() {
	Context("Test of roman-arabic equivalence of a given number", func() {
		It("check", func() {
			var romanequivalences = []struct {
				arabic int
				roman  string
			}{
				{arabic: 1, roman: "I"},
				{arabic: 2, roman: "II"},
				{arabic: 3, roman: "III"},
				{arabic: 4, roman: "IV"},
				{arabic: 5, roman: "V"},
				{arabic: 6, roman: "VI"},
				{arabic: 9, roman: "IX"},
				{arabic: 13, roman: "XIII"},
				{arabic: 15, roman: "XV"},
				{arabic: 19, roman: "XIX"},
				{arabic: 41, roman: "XLI"},
				{arabic: 50, roman: "L"},
				{arabic: 90, roman: "XC"},
				{arabic: 97, roman: "XCVII"},
				{arabic: 100, roman: "C"},
			}
			for _, arabicromanmap := range romanequivalences {
				Expect(Convertarabictoroman(arabicromanmap.arabic)).Should(BeEquivalentTo(arabicromanmap.roman))
			}
		})
	})
})
