package internal

import (
	"github.com/fraylopez/bank-go/internal/domain"
	"github.com/fraylopez/bank-go/internal/infrastructure/storage"
	"github.com/google/uuid"
	"testing"
)

func TestBank(t *testing.T) {
	t.Run("Open a Account", func(t *testing.T) {
		bank := GetBank()
		accountId, err := bank.OpenAccount("John Doe", "USD")
		if err != nil {
			t.Errorf("Error opening account: %v", err)
		}
		if uuid.Validate(accountId) != nil {
			t.Errorf("Invalid account id: %v", accountId)
		}
	})

	t.Run("Deposit to Account", func(t *testing.T) {
		bank := GetBank()
		accountId, _ := bank.OpenAccount("John Doe", "USD")
		err := bank.Deposit(accountId, 10, "USD")
		if err != nil {
			t.Errorf("Error depositing to account: %v", err)
		}
	})

	t.Run("Withdraw from Account", func(t *testing.T) {
		bank := GetBank()
		accountId, _ := bank.OpenAccount("John Doe", "USD")
		_ = bank.Deposit(accountId, 10, "USD")
		err := bank.Withdraw(accountId, 5, "USD")
		if err != nil {
			t.Errorf("Error withdrawing from account: %v", err)
		}
	})

	t.Run("Get balance from Account", func(t *testing.T) {
		bank := GetBank()
		accountId, _ := bank.OpenAccount("John Doe", "USD")
		_ = bank.Deposit(accountId, 10, "USD")
		balance, _ := bank.GetBalance(accountId)
		if !balance.Equals(domain.USD(10)) {
			t.Errorf("Expected balance to be 10, got %v", balance)
		}
	})
}

func GetBank() *Bank {
	return NewBank(storage.NewInMemoryAccountRepository())
}
