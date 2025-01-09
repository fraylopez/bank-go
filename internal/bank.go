package internal

import "bank/internal/domain"

// Bank is a struct that represents a bank
type Bank struct {
	accounts map[string]*domain.Account
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]*domain.Account),
	}
}

func (b *Bank) OpenAccount() (string, error) {
	account := domain.NewAccount()
	b.accounts[account.Id] = account
	return account.Id, nil
}

func (b *Bank) Deposit(accountId string, amount float64) error {
	account, err := getAccountById(b, accountId)
	if err != nil {
		return err
	}
	account.Deposit(amount)
	return nil
}

func (b *Bank) Withdraw(accountId string, amount float64) error {
	account, err := getAccountById(b, accountId)
	if err != nil {
		return err
	}
	return account.Withdraw(amount)
}

func getAccountById(bank *Bank, accountId string) (*domain.Account, error) {
	account := bank.accounts[accountId]
	if account == nil {
		return nil, &domain.AccountNotFoundError{}
	}
	return account, nil
}
