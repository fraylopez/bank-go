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
		account := BuildUSDAccount()
		if err := account.Deposit(USD(10)); err != nil {
			t.Errorf("Error depositing to account: %v", err)
		}
		if !account.Balance.Equals(USD(10)) {
			t.Errorf("Deposit did not add to the balance")
		}
	})

	t.Run("Withdraw subtracts from the balance", func(t *testing.T) {
		account := BuildUSDAccount()
		Deposit(t, account, USD(10))
		if err := account.Withdraw(USD(5)); err != nil {

			t.Errorf("Error withdrawing from account: %v", err)
		}

		if !account.Balance.Equals(USD(5)) {
			t.Errorf("Withdraw did not subtract from the balance")
		}
	})

	t.Run("Withdraw does not allow negative balance", func(t *testing.T) {
		account := BuildAccount()
		Deposit(t, account, USD(10))
		err := account.Withdraw(USD(15))
		assert.IsType(t, &NotEnoughFundsError{}, err)
	})

	t.Run("Prevents currency mismatch", func(t *testing.T) {
		account := BuildEURAccount()
		err := account.Deposit(USD(10))
		assert.IsType(t, &CurrencyMismatchError{}, err)
	})
}

func Deposit(t *testing.T, account *Account, amount Money) {
	err := account.Deposit(amount)
	assert.Nil(t, err)
}
