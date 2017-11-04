package bankaccount

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBankAccount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BankAccount")
}

var _ = Describe("BankAccount", func() {

	Context("Check credit and balnace", func() {
		It("Check credit and balance if deposited 1000 eur", func() {
			var balance float32 = 1000.0
			Expect(checkbalance()).To(Equal(balance))
		})

	})

})
