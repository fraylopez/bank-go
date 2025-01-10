package main

import (
	"fmt"
	"github.com/fraylopez/bank-go/internal"
	"github.com/fraylopez/bank-go/internal/infrastructure/http_api"
	"github.com/fraylopez/bank-go/internal/infrastructure/storage"
)

func main() {
	fmt.Println("Starting bank application...")
	accountRepository := storage.NewInMemoryAccountRepository()
	bank := internal.NewBank(accountRepository)
	http_api.Start(bank)
}
