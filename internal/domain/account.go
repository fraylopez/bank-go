package domain

import "github.com/google/uuid"

type Account struct {
	Id      string
	Holder  string
	Balance float64
}

func NewAccount(holder string) *Account {
	return &Account{
		Id:     uuid.NewString(),
		Holder: holder,
	}
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) error {
	if a.Balance < amount {
		return &NotEnoughFundsError{}
	}
	a.Balance -= amount
	return nil
}
