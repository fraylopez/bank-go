package account

import (
	"sync"

	"github.com/fraylopez/bank-go/internal/domain/money"
)

type TransferService interface {
	Transfer(fromAccountId string, toAccountId string, amount money.Money) error
}

type InMemoryTransferService struct {
	mutex             *sync.Mutex
	accountRepository AccountRepository
}

func NewTransferService(repository AccountRepository) TransferService {
	return &InMemoryTransferService{
		mutex:             new(sync.Mutex),
		accountRepository: repository,
	}

}

func (t *InMemoryTransferService) Transfer(fromAccountId string, toAccountId string, amount money.Money) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	fromAccount, err := t.accountRepository.GetAccountById(fromAccountId)
	if err != nil {
		return err
	}
	toAccount, err := t.accountRepository.GetAccountById(toAccountId)
	if err != nil {
		return err
	}
	if err := fromAccount.Withdraw(amount, nil); err != nil {
		return err
	}
	if err := toAccount.Deposit(amount); err != nil {
		return err
	}
	if err := t.accountRepository.UpdateAccount(fromAccount); err != nil {
		return err
	}
	if err := t.accountRepository.UpdateAccount(toAccount); err != nil {
		return err
	}
	return nil
}
