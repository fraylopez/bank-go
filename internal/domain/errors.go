package domain

type NotEnoughFundsError struct{}
type AccountNotFoundError struct{}

func (e *NotEnoughFundsError) Error() string {
	return "not enough funds"
}

func (e *AccountNotFoundError) Error() string {
	return "account not found"
}
