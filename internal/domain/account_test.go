package domain

import "testing"

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

}
