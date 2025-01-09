package domain

import "github.com/google/uuid"

type Account struct {
	Id      string
	Balance float64
}

func NewAccount() *Account {
	return &Account{
		Id: uuid.NewString(),
	}
}
