package storage_test

import (
	"testing"

	"github.com/fraylopez/bank-go/internal/domain/account"
	"github.com/fraylopez/bank-go/internal/infrastructure/storage"
)

func TestAccountRepository(t *testing.T) {
	tests := []account.AccountRepository{
		storage.NewInMemoryAccountRepository(),
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
	}
}
