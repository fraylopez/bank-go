package account_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraylopez/bank-go/internal/domain"
	"github.com/fraylopez/bank-go/internal/domain/account"
	"github.com/fraylopez/bank-go/internal/domain/loans"
	"github.com/fraylopez/bank-go/internal/domain/money"
)

func TestAccount(t *testing.T) {

	t.Run("Account has no balance on opening", func(t *testing.T) {
		acc := account.BuildAccount()
		assert.True(t, acc.Balance.Equals(money.USD(0)))
	})

	t.Run("deposit adds to the balance", func(t *testing.T) {
		acc := account.BuildUSDAccount()
		_ = acc.Deposit(money.USD(10))
		assert.True(t, acc.Balance.Equals(money.USD(10)))
	})

	t.Run("Withdraw subtracts from the balance", func(t *testing.T) {
		acc := account.BuildUSDAccount()
		_ = acc.Deposit(money.USD(10))
		_ = acc.Withdraw(money.USD(5), nil)
		assert.True(t, acc.Balance.Equals(money.USD(5)))
	})

	t.Run("Withdraw does not allow negative balance", func(t *testing.T) {
		acc := account.BuildUSDAccount()
		_ = acc.Deposit(money.USD(10))
		err := acc.Withdraw(money.USD(15), loans.NewNoLoansService())
		assert.IsType(t, &domain.NotEnoughFundsError{}, err)
	})

	t.Run("Prevents currency mismatch", func(t *testing.T) {
		acc := account.BuildEURAccount()
		err := acc.Deposit(money.USD(10))
		assert.IsType(t, &domain.CurrencyMismatchError{}, err)
	})
}
