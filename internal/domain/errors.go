package domain

type NotEnoughFundsError struct{}

func (e *NotEnoughFundsError) Error() string {
	return "not enough funds"
}
