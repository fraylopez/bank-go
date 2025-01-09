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

	t.Run("Deposit to Account", func(t *testing.T) {
		bank := NewBank()
		accountId, _ := bank.OpenAccount()
		err := bank.Deposit(accountId, 10)
		if err != nil {
			t.Errorf("Error depositing to account: %v", err)
		}
	})

	t.Run("Withdraw from Account", func(t *testing.T) {
		bank := NewBank()
		accountId, _ := bank.OpenAccount()
		_ = bank.Deposit(accountId, 10)
		err := bank.Withdraw(accountId, 5)
		if err != nil {
			t.Errorf("Error withdrawing from account: %v", err)
		}
	})

}
