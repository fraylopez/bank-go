package account

type AccountRepository interface {
	GetAccountById(id string) (*Account, error)
	OpenAccount(account *Account) error
	UpdateAccount(account *Account) error
}
