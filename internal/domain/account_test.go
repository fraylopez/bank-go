package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount(t *testing.T) {

	t.Run("Account has no balance on opening", func(t *testing.T) {
		account := NewAccount()
		if account.Balance != 0 {
			t.Errorf("Account balance is not zero")
		}
	})

	t.Run("Deposit adds to the balance", func(t *testing.T) {
		account := NewAccount()
		account.Deposit(10)
		if account.Balance != 10 {
			t.Errorf("Deposit did not add to the balance")
		}
	})

	t.Run("Withdraw subtracts from the balance", func(t *testing.T) {
		account := NewAccount()
		account.Deposit(10)
		if err := account.Withdraw(5); err != nil {
			t.Errorf("Error withdrawing from account: %v", err)
		}
		if account.Balance != 5 {
			t.Errorf("Withdraw did not subtract from the balance")
		}
	})

	t.Run("Withdraw does not allow negative balance", func(t *testing.T) {
		account := NewAccount()
		account.Deposit(10)
		err := account.Withdraw(15)
		assert.IsType(t, &NotEnoughFundsError{}, err)
	})

}
