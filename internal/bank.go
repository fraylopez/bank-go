package internal

import (
	"github.com/fraylopez/bank-go/internal/domain/account"
	"github.com/fraylopez/bank-go/internal/domain/money"
)

// Bank is a struct that represents a bank
type Bank struct {
	accountRepository account.AccountRepository
}

func NewBank(accountRepository account.AccountRepository) *Bank {
	return &Bank{
		accountRepository: accountRepository,
	}
}

func (b *Bank) OpenAccount(holder string, currency string) (string, error) {
	account := account.NewAccount(holder, currency)
	if err := b.accountRepository.OpenAccount(account); err != nil {
		return "", err
	}
	return account.Id, nil
}

func (b *Bank) Deposit(accountId string, amount float64, currency string) error {
	account, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return err
	}
	return account.Deposit(money.MoneyFrom(amount, currency))
}

func (b *Bank) Withdraw(accountId string, amount float64, currency string) error {
	account, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return err
	}
	return account.Withdraw(money.MoneyFrom(amount, currency))
}

func (b *Bank) GetBalance(accountId string) (money.Money, error) {
	account, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return money.Money{}, err
	}
	return account.Balance, nil
}

func (b *Bank) Transfer(fromAccountId string, toAccountId string, amount float64, currency string) error {
	fromAccount, err := b.accountRepository.GetAccountById(fromAccountId)
	if err != nil {
		return err
	}
	toAccount, err := b.accountRepository.GetAccountById(toAccountId)
	if err != nil {
		return err
	}
	if err := fromAccount.Withdraw(money.MoneyFrom(amount, currency)); err != nil {
		return err
	}
	return toAccount.Deposit(money.MoneyFrom(amount, currency))
}
