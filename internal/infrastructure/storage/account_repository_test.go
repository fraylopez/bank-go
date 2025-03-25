package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraylopez/bank-go/internal/domain/account"
	"github.com/fraylopez/bank-go/internal/domain/money"
	"github.com/fraylopez/bank-go/internal/infrastructure/storage"
)

func TestAccountRepository(t *testing.T) {
	tests := []account.AccountRepository{
		storage.NewInMemoryAccountRepository(),
		storage.NewTextFileAccountRepository("test_accounts.txt"),
	}

	for _, impl := range tests {
		t.Run("Open Account", func(t *testing.T) {
			account := account.BuildAccount()
			if err := impl.OpenAccount(account); err != nil {
				t.Errorf("Error opening account: %v", err)
			}
		})

		t.Run("Get Account by Id", func(t *testing.T) {
			account := account.BuildAccount()
			if err := impl.OpenAccount(account); err != nil {
				t.Errorf("Error opening account: %v", err)
			}
			_, err := impl.GetAccountById(account.Id)
			if err != nil {
				t.Errorf("Error getting account by id: %v", err)
			}
		})

		t.Run("Updates an account", func(t *testing.T) {
			acc := account.BuildAccount()
			_ = impl.OpenAccount(acc)
			_ = acc.Deposit(money.USD(100))

			_ = impl.UpdateAccount(acc)

			updatedAccount, _ := impl.GetAccountById(acc.Id)
			assert.True(t, updatedAccount.Balance.Equals(money.USD(100)))

		})
	}
}
