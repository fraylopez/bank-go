package internal

import (
	"sync"

	"github.com/fraylopez/bank-go/internal/domain/account"
	"github.com/fraylopez/bank-go/internal/domain/money"
)

// Bank is a struct that represents a bank
type Bank struct {
	accountRepository account.AccountRepository
	mutex             *sync.Mutex
}

func NewBank(accountRepository account.AccountRepository) *Bank {
	return &Bank{
		accountRepository: accountRepository,
		mutex:             new(sync.Mutex),
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
	err = acc.Withdraw(money.MoneyFrom(amount, currency))
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
	b.mutex.Lock()
	defer b.mutex.Unlock()

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
	if err := toAccount.Deposit(money.MoneyFrom(amount, currency)); err != nil {
		return err
	}
	if err := b.accountRepository.UpdateAccount(fromAccount); err != nil {
		return err
	}
	if err := b.accountRepository.UpdateAccount(toAccount); err != nil {
		return err
	}
	return nil
}
