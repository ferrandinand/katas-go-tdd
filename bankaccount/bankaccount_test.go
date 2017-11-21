package bankaccount

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBankAccount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BankAccount")
}

var _ = Describe("BankAccount", func() {
	AfterEach(func() {
		resetBalance()
	})

	Context("Check credit and balance", func() {
		It("Check credit and balance if we have no deposit", func() {
			Expect(balance()).To(Equal(0.0))
		})

		It("Check credit and balance if deposited 1000 eur", func() {
			deposit(1000.0)
			Expect(balance()).To(Equal(1000.0))
		})

		It("Check credit and balance if deposited 1500 eur", func() {
			deposit(1000.0)
			deposit(500.0)
			Expect(balance()).To(Equal(1500.0))
		})

	})
	Context("Check output", func() {
		It("Contains deposit data and balance", func() {
			deposit(500.0)
			currentdate := time.Now().Format("2/01/2006")
			Expect(outputMovement()).To(Equal("date || credit || balance\n" + currentdate + " || 500.00 || 500.00"))
		})
	})
})
