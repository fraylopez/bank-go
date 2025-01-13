package http_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fraylopez/bank-go/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	ts := httptest.NewServer(Handler(internal.BuildBank()))
	defer ts.Close()

	t.Run("should pass healthcheck", func(t *testing.T) {
		res, err := http.Get(ts.URL + "/health")

		if err != nil {
			t.Errorf("Error making request: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", res.StatusCode)
		}
	})

	t.Run("should create a new account", func(t *testing.T) {
		var requestBody = struct {
			Holder   string `json:"holder"`
			Currency string `json:"currency"`
		}{
			Holder:   "John Doe",
			Currency: "USD",
		}

		reqBody, err := json.Marshal(requestBody)

		res, err := http.Post(ts.URL+"/accounts", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Errorf("Error making request: %v", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", res.StatusCode)
		}

		var responseBody struct {
			AccountId string `json:"account_id"`
		}
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		if err != nil {
			t.Errorf("Error decoding response body: %v", err)
		}
		if responseBody.AccountId == "" {
			t.Errorf("Expected account id to be not empty")
		}
	})

	t.Run("should add funds", func(t *testing.T) {
		accountId := createAccount(ts)

		var requestBody = struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		}{
			Amount:   10,
			Currency: "USD",
		}

		reqBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Errorf("Error marshalling request body: %v", err)
		}

		res, err := http.Post(
			ts.URL+"/accounts/"+accountId+"/deposit",
			"application/json",
			bytes.NewBuffer(reqBody),
		)
		if err != nil {
			t.Errorf("Error making request: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", res.StatusCode)
		}
	})

	t.Run("should withdraw funds", func(t *testing.T) {
		accountId := createAccount(ts)
		deposit(ts, accountId, 10)

		var requestBody = struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		}{
			Amount:   5,
			Currency: "USD",
		}

		reqBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Errorf("Error marshalling request body: %v", err)
		}

		res, err := http.Post(
			ts.URL+"/accounts/"+accountId+"/withdraw",
			"application/json",
			bytes.NewBuffer(reqBody),
		)

		if err != nil {
			t.Errorf("Error making request: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", res.StatusCode)
		}

	})

	t.Run("should get balance", func(t *testing.T) {
		accountId := createAccount(ts)
		deposit(ts, accountId, 10)

		res, err := http.Get(ts.URL + "/accounts/" + accountId + "/balance")
		if err != nil {
			t.Errorf("Error making request: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", res.StatusCode)
		}

		var responseBody struct {
			Balance  float64 `json:"balance"`
			Currency string  `json:"currency"`
		}
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		if err != nil {
			t.Errorf("Error decoding response body: %v", err)
		}
		if responseBody.Balance != 10 || responseBody.Currency != "USD" {
			t.Errorf("Expected balance to be 10 USD, got %v %s", responseBody.Balance, responseBody.Currency)
		}
	})
}

func createAccount(ts *httptest.Server) string {
	var requestBody = struct {
		Holder   string `json:"holder"`
		Currency string `json:"currency"`
	}{
		Holder:   "John Doe",
		Currency: "USD",
	}

	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		panic(fmt.Sprintf("Error marshalling request body: %v", err))
	}
	res, err := http.Post(ts.URL+"/accounts", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(fmt.Sprintf("Error making request: %v", err))
	}

	var responseBody struct {
		AccountId string `json:"account_id"`
	}
	err = json.NewDecoder(res.Body).Decode(&responseBody)
	if err != nil {
		panic(fmt.Sprintf("Error decoding response body: %v", err))

	}
	return responseBody.AccountId
}

func deposit(ts *httptest.Server, accountId string, amount float64) {
	var requestBody = struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	}{
		Amount:   amount,
		Currency: "USD",
	}

	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		panic(fmt.Sprintf("Error marshalling request body: %v", err))
	}

	res, err := http.Post(
		ts.URL+"/accounts/"+accountId+"/deposit",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		panic(fmt.Sprintf("Error making request: %v", err))
	}
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Expected status code 200, got %v", res.StatusCode))
	}
}
