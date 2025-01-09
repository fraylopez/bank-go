package domain

import "github.com/google/uuid"

type Account struct {
	Id string
}

func NewAccount() *Account {
	return &Account{
		Id: uuid.NewString(),
	}
}
