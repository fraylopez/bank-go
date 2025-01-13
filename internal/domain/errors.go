package domain

type NotEnoughFundsError struct{}
type AccountNotFoundError struct{}
type CurrencyMismatchError struct{}

func (e *NotEnoughFundsError) Error() string {
	return "not enough funds"
}

func (e *AccountNotFoundError) Error() string {
	return "account not found"
}

func (e *CurrencyMismatchError) Error() string {
	return "currency mismatch"
}
