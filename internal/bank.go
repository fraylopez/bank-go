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

func (b *Bank) OpenAccount(holder string, currency string) (string, error) {
	account := domain.NewAccount(holder, currency)
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
	return account.Deposit(domain.MoneyFrom(amount, currency))
}

func (b *Bank) Withdraw(accountId string, amount float64, currency string) error {
	account, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return err
	}
	return account.Withdraw(domain.MoneyFrom(amount, currency))
}

func (b *Bank) GetBalance(accountId string) (domain.Money, error) {
	account, err := b.accountRepository.GetAccountById(accountId)
	if err != nil {
		return domain.Money{}, err
	}
	return account.Balance, nil
}
