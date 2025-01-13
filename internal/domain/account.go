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

func (a *Account) Deposit(amount float64) {
	a.Balance.Amount += amount
}

func (a *Account) Withdraw(amount float64) error {
	if a.Balance.Amount < amount {
		return &NotEnoughFundsError{}
	}
	a.Balance.Amount -= amount
	return nil
}
