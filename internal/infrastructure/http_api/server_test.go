package http_api

import (
	"encoding/json"
	"github.com/fraylopez/bank-go/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	ts := httptest.NewServer(Handler(internal.BuildBank()))
	defer ts.Close()

	t.Run("GET /health Should return 200 OK", func(t *testing.T) {
		res, err := http.Get(ts.URL + "/health")

		if err != nil {
			t.Errorf("Error making request: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", res.StatusCode)
		}
	})

	t.Run("POST /accounts should create a new account", func(t *testing.T) {
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
}
