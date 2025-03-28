package account

import (
	"github.com/google/uuid"

	"github.com/fraylopez/bank-go/internal/domain"
	"github.com/fraylopez/bank-go/internal/domain/loans"
	"github.com/fraylopez/bank-go/internal/domain/money"
)

type Account struct {
	Id       string
	Holder   string
	Currency money.Currencies
	Balance  money.Money
}

func NewAccount(holder string, currency string) *Account {
	return &Account{
		Id:       uuid.NewString(),
		Holder:   holder,
		Currency: money.Currency(currency),
		Balance:  money.NewMoney(currency),
	}
}

func (a *Account) Deposit(d money.Money) error {
	if newBalance, err := a.Balance.Add(d); err == nil {
		a.Balance = newBalance
	} else {
		return err
	}
	return nil
}

func (a *Account) Withdraw(w money.Money, service loans.LoanService) error {
	if a.Balance.IsLessThan(w) {
		allowed, err := service.CanAssignLoan(a.Balance, w)
		if err != nil {
			return err
		}
		if newBalance, err := a.Balance.Subtract(w); allowed && err == nil {
			a.Balance = newBalance
		} else {
			return &domain.NotEnoughFundsError{}
		}
	}
	if newBalance, err := a.Balance.Subtract(w); err == nil {
		a.Balance = newBalance
	} else {
		return err
	}
	return nil
}
