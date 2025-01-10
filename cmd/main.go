package main

import (
	"github.com/fraylopez/bank-go/internal"
	"github.com/fraylopez/bank-go/internal/infrastructure/http_api"
	"github.com/fraylopez/bank-go/internal/infrastructure/storage"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Starting bank application...")
	accountRepository := storage.NewInMemoryAccountRepository()
	bank := internal.NewBank(accountRepository)
	address := "localhost:8080"
	go http_api.Start(bank, address)
	log.Println("Bank application started!")
	health(address)
}

func health(address string) {
	for {
		res, err := http.Get("http://" + address + "/health")
		isHealthy := res.StatusCode == http.StatusOK && err == nil
		if !isHealthy {
			log.Println("Bank application is not healthy, shutting down...")
			panic("Bank application is not healthy!")
		} else {
			log.Println("Bank application is healthy!")
		}
		time.Sleep(time.Second * 5)
	}
}
