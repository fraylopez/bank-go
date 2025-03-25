package storage

import (
	"github.com/fraylopez/bank-go/internal/domain"
	"github.com/fraylopez/bank-go/internal/domain/account"
)

type InMemoryAccountRepository struct {
	accounts map[string]*account.Account
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		accounts: make(map[string]*account.Account),
	}
}

func (r *InMemoryAccountRepository) GetAccountById(accountId string) (*account.Account, error) {
	account, ok := r.accounts[accountId]
	if !ok {
		return nil, &domain.AccountNotFoundError{}
	}
	return account, nil
}

func (r *InMemoryAccountRepository) OpenAccount(account *account.Account) error {
	r.accounts[account.Id] = account
	return nil
}

func (r *InMemoryAccountRepository) UpdateAccount(account *account.Account) error {
	if _, ok := r.accounts[account.Id]; !ok {
		return &domain.AccountNotFoundError{}
	}
	r.accounts[account.Id] = account
	return nil
}
