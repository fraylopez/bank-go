package domain

// Bank is a struct that represents a bank
type Bank struct {
	accounts map[string]float64
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]float64),
	}
}

func (b *Bank) OpenAccount() (string, error) {
	return NewAccount().Id, nil
}
