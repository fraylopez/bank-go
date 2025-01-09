package http_api

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	setup := func() func() {
		server := NewServer()
		go Start(server)
		return func() {
			Shutdown(server)
		}
	}

	t.Run("/health Should return 200 OK", func(t *testing.T) {
		teardown := setup()
		defer teardown()

		res, err := http.Get("http://localhost:8080/health")

		if err != nil {
			t.Errorf("Error making request: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", res.StatusCode)
		}
	})
}
