package internal

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/fraylopez/bank-go/internal/domain/money"
	"github.com/fraylopez/bank-go/internal/infrastructure/storage"
)

func TestBank_UseCases(t *testing.T) {
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
		if !balance.Equals(money.USD(10)) {
			t.Errorf("Expected balance to be 10, got %v", balance)
		}
	})

	t.Run("Transfer between accounts", func(t *testing.T) {
		bank := GetBank()
		accountId1, _ := bank.OpenAccount("John Doe", "USD")
		accountId2, _ := bank.OpenAccount("Jane Doe", "USD")
		_ = bank.Deposit(accountId1, 10, "USD")
		err := bank.Transfer(accountId1, accountId2, 5, "USD")
		if err != nil {
			t.Errorf("Error transferring between accounts: %v", err)
		}
		balance1, _ := bank.GetBalance(accountId1)
		balance2, _ := bank.GetBalance(accountId2)

		assert.True(t, balance1.Equals(money.USD(5)))
		assert.True(t, balance2.Equals(money.USD(5)))
	})
}

func GetBank() *Bank {
	return NewBank(storage.NewTextFileAccountRepository("test_accounts.txt"))
}
