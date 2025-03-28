package account

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraylopez/bank-go/internal/domain"
	"github.com/fraylopez/bank-go/internal/domain/money"
)

func TestAccount(t *testing.T) {

	t.Run("Account has no balance on opening", func(t *testing.T) {
		account := BuildAccount()
		assert.True(t, account.Balance.Equals(money.USD(0)))
	})

	t.Run("deposit adds to the balance", func(t *testing.T) {
		account := BuildUSDAccount()
		_ = account.Deposit(money.USD(10))
		assert.True(t, account.Balance.Equals(money.USD(10)))
	})

	t.Run("Withdraw subtracts from the balance", func(t *testing.T) {
		account := BuildUSDAccount()
		_ = account.Deposit(money.USD(10))
		_ = account.Withdraw(money.USD(5))
		assert.True(t, account.Balance.Equals(money.USD(5)))
	})

	t.Run("Withdraw does not allow negative balance", func(t *testing.T) {
		account := BuildUSDAccount()
		_ = account.Deposit(money.USD(10))
		err := account.Withdraw(money.USD(15))
		assert.IsType(t, &domain.NotEnoughFundsError{}, err)
	})

	t.Run("Prevents currency mismatch", func(t *testing.T) {
		account := BuildEURAccount()
		err := account.Deposit(money.USD(10))
		assert.IsType(t, &domain.CurrencyMismatchError{}, err)
	})
}
