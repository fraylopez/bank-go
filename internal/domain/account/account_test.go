package account

import (
	"github.com/fraylopez/bank-go/internal/domain"
	"github.com/fraylopez/bank-go/internal/domain/money"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAccount(t *testing.T) {

	t.Run("Account has no balance on opening", func(t *testing.T) {
		account := BuildAccount()
		if !reflect.DeepEqual(account.Balance, money.USD(0)) {
			t.Errorf("Account balance is not zero")
		}
	})

	t.Run("deposit adds to the balance", func(t *testing.T) {
		account := BuildUSDAccount()
		if err := account.Deposit(money.USD(10)); err != nil {
			t.Errorf("Error depositing to account: %v", err)
		}
		if !account.Balance.Equals(money.USD(10)) {
			t.Errorf("deposit did not add to the balance")
		}
	})

	t.Run("Withdraw subtracts from the balance", func(t *testing.T) {
		account := BuildUSDAccount()
		deposit(t, account, money.USD(10))
		if err := account.Withdraw(money.USD(5)); err != nil {

			t.Errorf("Error withdrawing from account: %v", err)
		}

		if !account.Balance.Equals(money.USD(5)) {
			t.Errorf("Withdraw did not subtract from the balance")
		}
	})

	t.Run("Withdraw does not allow negative balance", func(t *testing.T) {
		account := BuildAccount()
		deposit(t, account, money.USD(10))
		err := account.Withdraw(money.USD(15))
		assert.IsType(t, &domain.NotEnoughFundsError{}, err)
	})

	t.Run("Prevents currency mismatch", func(t *testing.T) {
		account := BuildEURAccount()
		err := account.Deposit(money.USD(10))
		assert.IsType(t, &domain.CurrencyMismatchError{}, err)
	})
}

func deposit(t *testing.T, account *Account, amount money.Money) {
	err := account.Deposit(amount)
	assert.Nil(t, err)
}
