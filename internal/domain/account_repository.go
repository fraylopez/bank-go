package domain

type AccountRepository interface {
	GetAccountById(id string) (*Account, error)
	OpenAccount(account *Account) error
}
