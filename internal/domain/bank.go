package domain

// Bank is a struct that represents a bank
type Bank struct {
	accounts map[string]*Account
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]*Account),
	}
}

func (b *Bank) OpenAccount() (string, error) {
	account := NewAccount()
	b.accounts[account.Id] = account
	return account.Id, nil
}

func (b *Bank) Deposit(accountId string, amount float64) error {
	account := b.accounts[accountId]
	if account == nil {
		return &AccountNotFoundError{}
	}
	account.Deposit(amount)
	return nil
}

func (b *Bank) Withdraw(accountId string, amount float64) error {
	account := b.accounts[accountId]
	if account == nil {
		return &AccountNotFoundError{}
	}
	return account.Withdraw(amount)
}
