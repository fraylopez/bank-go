package storage

import "github.com/fraylopez/bank-go/internal/domain"

type InMemoryAccountRepository struct {
	accounts map[string]*domain.Account
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		accounts: make(map[string]*domain.Account),
	}
}

func (r *InMemoryAccountRepository) GetAccountById(accountId string) (*domain.Account, error) {
	account, ok := r.accounts[accountId]
	if !ok {
		return nil, &domain.AccountNotFoundError{}
	}
	return account, nil
}

func (r *InMemoryAccountRepository) OpenAccount(account *domain.Account) error {
	r.accounts[account.Id] = account
	return nil
}
