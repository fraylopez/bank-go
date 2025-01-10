package storage

import (
	"github.com/fraylopez/bank-go/internal/domain"
	"testing"
)

func TestAccountRepository(t *testing.T) {
	tests := []domain.AccountRepository{
		NewInMemoryAccountRepository(),
	}

	for _, impl := range tests {
		t.Run("Open Account", func(t *testing.T) {
			account := domain.BuildAccount()
			if err := impl.OpenAccount(account); err != nil {
				t.Errorf("Error opening account: %v", err)
			}
		})

		t.Run("Get Account by Id", func(t *testing.T) {
			account := domain.BuildAccount()
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
