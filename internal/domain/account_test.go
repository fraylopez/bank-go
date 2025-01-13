package domain

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAccount(t *testing.T) {

	t.Run("Account has no balance on opening", func(t *testing.T) {
		account := BuildAccount()
		if !reflect.DeepEqual(account.Balance, USD(0)) {
			t.Errorf("Account balance is not zero")
		}
	})

	t.Run("Deposit adds to the balance", func(t *testing.T) {
		account := BuildAccount()
		account.Deposit(10)
		if !account.Balance.Equals(USD(10)) {
			t.Errorf("Deposit did not add to the balance")
		}
	})

	t.Run("Withdraw subtracts from the balance", func(t *testing.T) {
		account := BuildAccount()
		account.Deposit(10)
		if err := account.Withdraw(5); err != nil {
			t.Errorf("Error withdrawing from account: %v", err)
		}
		if !account.Balance.Equals(USD(5)) {
			t.Errorf("Withdraw did not subtract from the balance")
		}
	})

	t.Run("Withdraw does not allow negative balance", func(t *testing.T) {
		account := BuildAccount()
		account.Deposit(10)
		err := account.Withdraw(15)
		assert.IsType(t, &NotEnoughFundsError{}, err)
	})
}
