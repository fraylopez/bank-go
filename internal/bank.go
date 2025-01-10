package internal

import "github.com/fraylopez/bank-go/internal/domain"

// Bank is a struct that represents a bank
type Bank struct {
	accounts          map[string]*domain.Account
	accountRepository domain.AccountRepository
}

func NewBank(accountRepository domain.AccountRepository) *Bank {
	return &Bank{
		accounts:          make(map[string]*domain.Account),
		accountRepository: accountRepository,
	}
}

func (b *Bank) OpenAccount() (string, error) {
	account := domain.NewAccount()
	if err := b.accountRepository.OpenAccount(account); err != nil {
		return "", err
	}
	return account.Id, nil
}

func (b *Bank) Deposit(accountId string, amount float64) error {
	account, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return err
	}
	account.Deposit(amount)
	return nil
}

func (b *Bank) Withdraw(accountId string, amount float64) error {
	account, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return err
	}
	return account.Withdraw(amount)
}
