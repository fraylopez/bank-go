package internal

import (
	"github.com/fraylopez/bank-go/internal/domain/account"
	"github.com/fraylopez/bank-go/internal/domain/loans"
	"github.com/fraylopez/bank-go/internal/domain/money"
)

// Bank is a struct that represents a bank
type Bank struct {
	accountRepository account.AccountRepository
	transferService   account.TransferService
	loanService       loans.LoanService
}

func NewBank(accountRepository account.AccountRepository) *Bank {
	return &Bank{
		accountRepository: accountRepository,
		transferService:   account.NewTransferService(accountRepository),
		loanService:       loans.NewNoLoansService(),
	}
}

func (b *Bank) OpenAccount(holder string, currency string) (string, error) {
	newAccount := account.NewAccount(holder, currency)
	if err := b.accountRepository.OpenAccount(newAccount); err != nil {
		return "", err
	}
	return newAccount.Id, nil
}

func (b *Bank) Deposit(accountId string, amount float64, currency string) error {
	acc, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return err
	}
	err = acc.Deposit(money.MoneyFrom(amount, currency))
	if err != nil {
		return err
	}
	return b.accountRepository.UpdateAccount(acc)
}

func (b *Bank) Withdraw(accountId string, amount float64, currency string) error {
	acc, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return err
	}
	err = acc.Withdraw(money.MoneyFrom(amount, currency), b.loanService)
	if err != nil {
		return err
	}
	return b.accountRepository.UpdateAccount(acc)
}

func (b *Bank) GetBalance(accountId string) (money.Money, error) {
	acc, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return money.Money{}, err
	}
	return acc.Balance, nil
}

func (b *Bank) Transfer(fromAccountId string, toAccountId string, amount float64, currency string) error {
	return b.transferService.Transfer(fromAccountId, toAccountId, money.MoneyFrom(amount, currency))
}
