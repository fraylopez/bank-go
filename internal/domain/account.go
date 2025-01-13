package domain

import "github.com/google/uuid"

type Account struct {
	Id       string
	Holder   string
	Currency string
	Balance  Money
}

func NewAccount(holder string, currency string) *Account {
	return &Account{
		Id:       uuid.NewString(),
		Holder:   holder,
		Currency: currency,
		Balance:  NewMoney(currency),
	}
}

func (a *Account) Deposit(d Money) error {
	if newBalance, err := a.Balance.Add(d); err == nil {
		a.Balance = newBalance
	} else {
		return err
	}
	return nil
}

func (a *Account) Withdraw(w Money) error {
	if a.Balance.Amount < w.Amount {
		return &NotEnoughFundsError{}
	}
	if newBalance, err := a.Balance.Subtract(w); err == nil {
		a.Balance = newBalance
	} else {
		return err
	}
	return nil
}
