package domain

import (
	"github.com/google/uuid"
	"testing"
)

func TestBank(t *testing.T) {
	t.Run("Open a Account", func(t *testing.T) {
		bank := NewBank()
		accountId, err := bank.OpenAccount()
		if err != nil {
			t.Errorf("Error opening account: %v", err)
		}
		if uuid.Validate(accountId) != nil {
			t.Errorf("Invalid account id: %v", accountId)
		}
	})

}
