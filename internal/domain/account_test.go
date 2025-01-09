package domain

import "testing"

func TestAccount(t *testing.T) {

	t.Run("Account can be instantiated", func(t *testing.T) {
		account := NewAccount()
		if account == nil {
			t.Errorf("Account was not instantiated")
		}
	})
}
