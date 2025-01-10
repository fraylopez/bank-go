package internal

import (
	"github.com/fraylopez/bank-go/internal/infrastructure/storage"
)

func BuildBank() *Bank {
	repo := storage.NewInMemoryAccountRepository()
	return NewBank(repo)
}
