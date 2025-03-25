package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraylopez/bank-go/internal/domain/money"
	"github.com/fraylopez/bank-go/internal/infrastructure/storage"
)

func TestBankTextFileRepoIntegration(t *testing.T) {
	t.Run("should create and retrieve an account", func(t *testing.T) {
		bank := GetTestBank()
		accountID, _ := bank.OpenAccount("John Doe", "USD")
		_ = bank.Deposit(accountID, 1000, "USD")

		theSameBank := GetTestBank()
		balance, _ := theSameBank.GetBalance(accountID)

		assert.True(t, balance.Equals(money.USD(1000)))
	})
}

func GetTestBank() *Bank {
	return &Bank{
		accountRepository: storage.NewTextFileAccountRepository("test_accounts.txt"),
	}
}
