package account_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraylopez/bank-go/internal/domain/account"
	"github.com/fraylopez/bank-go/internal/domain/money"
	"github.com/fraylopez/bank-go/internal/infrastructure/storage"
)

func TestTransferService(t *testing.T) {
	// Test cases for TransferService methods
	t.Run("should transfer balance to account", func(t *testing.T) {
		repo := storage.NewTextFileAccountRepository("test_accounts.txt")

		acc1 := account.BuildEURAccount()
		acc2 := account.BuildEURAccount()
		_ = acc1.Deposit(money.EUR(10))

		_ = repo.OpenAccount(acc1)
		_ = repo.OpenAccount(acc2)

		transferService := account.NewTransferService(repo)

		err := transferService.Transfer(acc1.Id, acc2.Id, money.EUR(5))

		updatedAcc1, _ := repo.GetAccountById(acc1.Id)
		updatedAcc2, _ := repo.GetAccountById(acc2.Id)

		assert.Nil(t, err)

		assert.Equal(t, 5.0, updatedAcc1.Balance.Amount)
		assert.Equal(t, 5.0, updatedAcc2.Balance.Amount)
	})

	t.Run("should transfer concurrently", func(t *testing.T) {
		repo := storage.NewTextFileAccountRepository("test_accounts.txt")

		acc1 := account.BuildEURAccount()
		acc2 := account.BuildEURAccount()
		_ = acc1.Deposit(money.EUR(10))

		_ = repo.OpenAccount(acc1)
		_ = repo.OpenAccount(acc2)

		transferService := account.NewTransferService(repo)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			err := transferService.Transfer(acc1.Id, acc2.Id, money.EUR(5))
			assert.Nil(t, err)
		}()

		go func() {
			defer wg.Done()
			err := transferService.Transfer(acc1.Id, acc2.Id, money.EUR(3))
			assert.Nil(t, err)
		}()

		wg.Wait()

		updatedAcc1, _ := repo.GetAccountById(acc1.Id)
		updatedAcc2, _ := repo.GetAccountById(acc2.Id)

		assert.Equal(t, 2.0, updatedAcc1.Balance.Amount)
		assert.Equal(t, 8.0, updatedAcc2.Balance.Amount)
	})
}
