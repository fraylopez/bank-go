package domain

import "testing"

func TestAccount(t *testing.T) {

	t.Run("Account has no balance on opening", func(t *testing.T) {
		account := NewAccount()
		if account.Balance != 0 {
			t.Errorf("Account balance is not zero")
		}
	})
}
