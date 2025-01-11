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
		res, err := http.Post(ts.URL+"/accounts", "application/json", nil)
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
			Amount float64 `json:"amount"`
		}{
			Amount: 10,
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

}

func createAccount(ts *httptest.Server) string {
	res, err := http.Post(ts.URL+"/accounts", "application/json", nil)
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
